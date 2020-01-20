package components

import (
	"github.com/aliyunmq/mq-http-go-sdk"
	"github.com/gogap/errors"
	"go.uber.org/zap"
	"strings"
	"time"
)

var RMQ *Rmq

type Rmq struct {
	client   mq_http_sdk.MQClient
	Producer mq_http_sdk.MQProducer
	Consumer []mq_http_sdk.MQConsumer
}

type Handler interface {
	HandleMessage(m mq_http_sdk.ConsumeMessageEntry) error
}

type HandlerFunc func(m mq_http_sdk.ConsumeMessageEntry) error

func (h HandlerFunc) HandleMessage(m mq_http_sdk.ConsumeMessageEntry) error {
	return h(m)
}

type MQConfig struct {
	Endpoint      string
	AccessKey     string
	AccessSecret  string
	Waitseconds   int64
	NumOfMessages int32
}

func InitMq(cfg *MQConfig) error {
	rmq := &Rmq{}
	rmq.client = mq_http_sdk.NewAliyunMQClient(cfg.Endpoint, cfg.AccessKey, cfg.AccessSecret, "")
	RMQ = rmq

	return nil
}

func (rmq *Rmq) NewProducer(instanceId string, topic string) {
	//实例化生产者
	rmq.Producer = rmq.client.GetProducer(instanceId, topic)
}

func (rmq *Rmq) NewConsumer(handle Handler, instanceId string, topic string, groupId string, tag string) {
	//实例化消费者
	consumer := rmq.client.GetConsumer(instanceId, topic, groupId, tag)
	go rmq.OnMessage(consumer, handle)
	rmq.Consumer = append(rmq.Consumer, consumer)
}

func (rmq *Rmq) Publish(tag string, msg string) bool {
	ret, err := rmq.Producer.PublishMessage(mq_http_sdk.PublishMessageRequest{
		MessageBody: msg,
		MessageTag:  tag,
	})

	if err != nil {
		L.Error("rocketMQ", zap.String("publish", err.Error()))
		return false
	}

	L.Debug("rocketMQ_publish", zap.String("message", msg), zap.String("id", ret.MessageId))

	return true
}

func (rmq *Rmq) OnMessage(consumer mq_http_sdk.MQConsumer, handle Handler) {
	for {
		endChan := make(chan int)
		respChan := make(chan mq_http_sdk.ConsumeMessageResponse)
		errChan := make(chan error)
		go func() {
			select {
			case resp := <-respChan:
				{
					// 处理业务逻辑
					var handles []string
					for _, v := range resp.Messages {
						handles = append(handles, v.ReceiptHandle)
						err := handle.HandleMessage(v)
						if err != nil {
							L.Error("rocketMQ", zap.String("consume", err.Error()))
						}
					}

					// NextConsumeTime前若不确认消息消费成功，则消息会重复消费
					// 消息句柄有时间戳，同一条消息每次消费拿到的都不一样
					ackErr := consumer.AckMessage(handles)
					if ackErr != nil {
						// 某些消息的句柄可能超时了会导致确认不成功
						for _, errAckItem := range ackErr.(errors.ErrCode).Context()["Detail"].([]mq_http_sdk.ErrAckItem) {
							L.Error("rocketMQ ack", zap.String("handle", errAckItem.ErrorHandle),
								zap.String("code", errAckItem.ErrorCode),
								zap.String("msg", errAckItem.ErrorMsg))
						}
						time.Sleep(time.Duration(3) * time.Second)
					}

					endChan <- 1
				}
			case err := <-errChan:
				{
					if !strings.Contains(err.(errors.ErrCode).Error(), "MessageNotExist") {
						L.Error("rocketMQ", zap.String("consume", err.Error()))
						time.Sleep(time.Duration(3) * time.Second)
					}
					endChan <- 1
				}
			case <-time.After(35 * time.Second):
				{
					L.Error("timeout_of_consumer_message")
					endChan <- 1
				}
			}
		}()

		consumer.ConsumeMessage(respChan, errChan,
			E.Opts.MQConfig.NumOfMessages,
			E.Opts.MQConfig.Waitseconds,
		)
		<-endChan
	}
}
