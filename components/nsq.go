package components

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/joselee214/j7f/components/mq"
	"go.uber.org/zap"
)

var N *Nsq

type Nsq struct {
	Producer *mq.Producer
	Consumer []*mq.Consumer
}

type NsqConfig struct {
	Config   *mq.Config
	Producer string
	Consumer []string
}

func InitNsq(cfg *NsqConfig) error {
	n := &Nsq{}
	producer, err := mq.NewProducer(cfg.Config)
	if err != nil {
		return fmt.Errorf("new nsq producer %s", err)
	}
	if err := producer.ConnectToNSQLookupd(cfg.Producer, cfg.Config.PoolCap); err != nil {
		return err
	}
	n.Producer = producer
	N = n
	return nil
}

func (n *Nsq) NewConsumer(topic string, channel string, handler nsq.Handler) error {
	consumer, err := mq.NewConsumer(topic, channel, E.Opts.NsqConfig.Config)
	if err != nil {
		return err
	}

	//consumer.SetLogger(L, E.Opts.LogConfig.Level)
	consumer.AddConcurrentHandlers(handler, E.Opts.NsqConfig.Config.Concurrency)
	if err = consumer.ConnectToNSQLookupds(E.Opts.NsqConfig.Consumer); err != nil {
		return err
	}

	n.Consumer = append(n.Consumer, consumer)

	return nil
}

func (n *Nsq) Publish(topic string, msg []byte) bool {
	if err := n.Producer.Publish(topic, msg); err != nil {
		L.Error("nsq", zap.String("publish", err.Error()))
		return false
	}
	L.Debug("nsq_publish", zap.ByteString("message", msg))
	return true
}
