package components

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"sync"
	"time"
	"context"
)

const(
	MAX_CONNECTION = 10
	INITIAL_CONNECTION = 1
	AVAILABLE = false
	USED = true

)

var mu sync.RWMutex

/*
clientList: the client pool
clientAvailable: the available flag, means the location and available flag in the  client pool
size: the size of allocated client pool <= MAX_CONNECTION
*/
type mongodata struct{
	client *mongo.Client
	pos int
	flag bool

}

type ClientPool struct{
	clientList [MAX_CONNECTION]mongodata
	size int
}

var cp ClientPool

//initial the connection to the pool
func init(){
	for size := 0;  size < INITIAL_CONNECTION || size < MAX_CONNECTION; size++ {
		err := cp.allocateCToPool(size)
		//fmt.Print("init - initial create the connect pool failed, size: ",size ,err)
		if err != nil {
			L.Info("init - initial create the connect pool failed, size: ", zap.Int("size",size) , zap.Error(err) )
		}
	}
}

func Dbconnect() (client *mongo.Client, err error){
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
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

func Dbdisconnect(client *mongo.Client) (err error){
	err = client.Disconnect(context.TODO())
	if err != nil {
		//fmt.Print("Dbdisconnect - disconnect mongodb failed: ", err)
		L.Info("Dbdisconnect - disconnect mongodb failed: ", zap.Error(err) )
	}
	return err
}

//create a new database connection to the pool
func (cp *ClientPool) allocateCToPool(pos int) (err error){
	cp.clientList[pos].client, err = Dbconnect()
	if err != nil {
		//fmt.Print("allocateCToPool - allocateCToPool failed,position: ", pos, err)
		L.Info("allocateCToPool - allocateCToPool failed,position: ", zap.Int("pos",pos), zap.Error(err) )
		return err
	}

	cp.clientList[pos].flag = USED
	cp.clientList[pos].pos = pos
	return nil
}

//apply a connection from the pool
func (cp *ClientPool) getCToPool(pos int){
	cp.clientList[pos].flag = USED
}

//free a connection back to the pool
func (cp *ClientPool) putCBackPool(pos int){
	cp.clientList[pos].flag = AVAILABLE
}

//program apply a database connection
func GetClient() (mongoclient *mongodata,  err error) {
	mu.RLock()
	for i:=1; i<cp.size; i++ {
		if cp.clientList[i].flag == AVAILABLE{
			return &cp.clientList[i], nil
		}
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()
	if cp.size < MAX_CONNECTION{
		err = cp.allocateCToPool(cp.size)
		if err != nil {
			//fmt.Print("GetClient - DB pooling allocate failed: ", err)
			L.Info("GetClient - DB pooling allocate failed: ", zap.Error(err) )
			return nil, err
		}

		pos := cp.size
		cp.size++
		return &cp.clientList[pos], nil
	} else {
		//fmt.Print("GetClient - DB pooling is fulled: ", err)
		L.Info("GetClient - DB pooling is fulled: ", zap.Error(err) )
		return nil, errors.New("DB pooling is fulled")
	}
}

//program release a connection
func ReleaseClient(mongoclient *mongodata){
	mu.Lock()
	cp.putCBackPool(mongoclient.pos)
	mu.Unlock()
}