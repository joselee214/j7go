package components

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"sync"
	"time"
)

var mongoconfigset *MongoConfig

type MongoConfig struct {
	Addr string
	MaxIdle int
	MaxActive int
}

const(
	AVAILABLE = false
	USED = true
)

var mu sync.RWMutex

/*
clientList: the client pool
clientAvailable: the available flag, means the location and available flag in the  client pool
size: the size of allocated client pool <= MAX_CONNECTION
*/
type MongoClient struct{
	Client *mongo.Client
	pos int
	flag bool
}

//释放到链接池
func (mongoclient *MongoClient) Release() {
	mu.Lock()
	MongoPool.putCBackPool(mongoclient.pos)
	mu.Unlock()
}

type MongoClientPool struct{
	clientList []MongoClient
	size int
}

var MongoPool MongoClientPool

//从配置初始化...
func NewMongoPool(cfg *MongoConfig) {
	mongoconfigset = cfg
	MongoPool.clientList = make([]MongoClient,mongoconfigset.MaxActive)
	for size := 0;  size < mongoconfigset.MaxIdle && size < mongoconfigset.MaxActive ; size++ {

		err := MongoPool.allocateCToPool(size)

		if err != nil {
			L.Info("init - initial create the connect pool failed, size: ", zap.Int("size",size) , zap.Error(err) )
		}
	}
}

func MongoDbConnectSingleton() (client *mongo.Client, err error){

	client, err = mongo.NewClient(options.Client().ApplyURI(mongoconfigset.Addr) )
	if err != nil {
		//fmt.Print("Dbconnect - connect mongodb failed: ", err)
		L.Info("Dbconnect - connect mongodb failed: ", zap.Error(err) )

		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		//fmt.Print("Dbconnect - connect mongodb ctx failed: ", err)
		L.Info("Dbconnect - connect mongodb ctx failed: ", zap.Error(err) )
		return nil, err
	}
	return client,nil
}

//从链接池获取一个链接...
func MongoGetClient() (mongoclient *MongoClient,  err error) {
	mu.RLock()
	for i:=1; i<MongoPool.size; i++ {
		if MongoPool.clientList[i].flag == AVAILABLE{
			return &MongoPool.clientList[i], nil
		}
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()
	if MongoPool.size < mongoconfigset.MaxActive {
		err = MongoPool.allocateCToPool(MongoPool.size)
		if err != nil {
			//fmt.Print("GetClient - DB pooling allocate failed: ", err)
			L.Info("GetClient - DB pooling allocate failed: ", zap.Error(err) )
			return nil, err
		}

		pos := MongoPool.size
		MongoPool.size++
		return &MongoPool.clientList[pos], nil
	} else {
		//fmt.Print("GetClient - DB pooling is fulled: ", err)
		L.Info("GetClient - DB pooling is fulled: ", zap.Error(err) )
		return nil, errors.New("DB pooling is fulled")
	}
}



//func MongoDbDisconnect(client *mongo.Client) (err error){
//	err = client.Disconnect(context.TODO())
//	if err != nil {
//		//fmt.Print("Dbdisconnect - disconnect mongodb failed: ", err)
//		L.Info("Dbdisconnect - disconnect mongodb failed: ", zap.Error(err) )
//	}
//	return err
//}

//create a new database connection to the pool
func (cp *MongoClientPool) allocateCToPool(pos int) (err error){
	MongoPool.clientList[pos].Client, err = MongoDbConnectSingleton()

	if err != nil {
		L.Info("allocateCToPool - allocateCToPool failed,position: ", zap.Int("pos",pos), zap.Error(err) )
		return err
	}

	MongoPool.clientList[pos].flag = USED
	MongoPool.clientList[pos].pos = pos
	return nil
}

//apply a connection from the pool
func (cp *MongoClientPool) getCToPool(pos int){
	MongoPool.clientList[pos].flag = USED
}

//free a connection back to the pool
func (cp *MongoClientPool) putCBackPool(pos int){
	MongoPool.clientList[pos].flag = AVAILABLE
}
