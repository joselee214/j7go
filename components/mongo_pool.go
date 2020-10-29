package components

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"sync"
	"time"
)

var mongoconfigset *MongoConfig

type MongoConfig struct {
	Addr string
	MaxIdle int
	MaxActive int
	TimeOut int
}

const(
	AVAILABLE = 1
	USING = 2
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
	flag int
	idlets int64
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
	//for size := 0;  size < mongoconfigset.MaxActive ; size++ {
	//	err := MongoPool.allocateCToPool(size)
	//	if err != nil {
	//		L.Info("init - initial create the connect pool failed, size: ", zap.Int("size",size) , zap.Error(err) )
	//	}
	//}
	go scanPoolList()
}

func scanPoolList()  {
	for {
		time.Sleep( 5 * time.Second )
		mu.Lock()
		//fmt.Println(" 1 MongoPool",MongoPool)
		nowts := time.Now().UnixNano()
		for i:=0; i<MongoPool.size; i++ {
			if MongoPool.clientList[i].flag == AVAILABLE{
				if ( nowts - MongoPool.clientList[i].idlets ) > (int64(mongoconfigset.MaxIdle) * int64(time.Second)) {
					MongoPool.clientList[i].Client.Disconnect(context.Background())
					MongoPool.clientList[i] = MongoClient{}
				}
			}
		}
		mu.Unlock()
	}
}

func ResetMongoConfig(cfg *MongoConfig){
	if cfg != nil {
		if len(MongoPool.clientList) == 0 {
			NewMongoPool(cfg)
		} else {
			mongoconfigset = cfg
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
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoconfigset.TimeOut) )
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second )
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		//fmt.Print("Dbconnect - connect mongodb ctx failed: ", err)
		L.Info("Dbconnect - connect mongodb ctx failed: ", zap.Error(err) )
		return nil, err
	}
	//if err = pingError(client); err !=nil {
	//	return nil,err
	//}
	return client,nil
}

func pingError(client *MongoClient) error {
	if err := client.Client.Ping(context.Background(),readpref.Primary()); err !=nil {
		MongoPool.allocateCToPool(client.pos)
		return err
	}
	return  nil
}

//从链接池获取一个链接...
func MongoGetClient() (*MongoClient,error) {
	mgcli,err := getClient()
	if err == nil {
		err = pingError(mgcli)
		if err!=nil {
			return nil,err
		}
	}
	MongoPool.clientList[mgcli.pos].flag = USING
	return  mgcli,err
}

func getClient() (mongoclient *MongoClient,  err error) {
	mu.RLock()
	for i:=0; i<MongoPool.size; i++ {
		if MongoPool.clientList[i].flag == AVAILABLE{
			mu.RUnlock()
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

	MongoPool.clientList[pos].flag = AVAILABLE
	MongoPool.clientList[pos].pos = pos
	MongoPool.clientList[pos].idlets = time.Now().UnixNano()
	return nil
}


//free a connection back to the pool
func (cp *MongoClientPool) putCBackPool(pos int){
	MongoPool.clientList[pos].flag = AVAILABLE
	MongoPool.clientList[pos].idlets = time.Now().UnixNano()

}
