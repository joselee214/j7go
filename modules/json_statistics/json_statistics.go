package json_statistics



import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joselee214/j7f/components/http/server"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"compress/gzip"
	"j7go/components"
	"time"
	"context"
)

var dataChan chan []interface {}
var chanlen int = 2000000

var databaseconfig string = "json_statistics"
var collectiontableconfig string = "statistics"

func Init(g *gin.Engine) {

	dataChan = make(chan []interface {},chanlen)

	s := &JsonStatisticsController{}
	g.GET("/favicon.ico",s.noop )//注册接口
	g.POST("/c",s.c)
	g.GET("/stat",s.stat )//注册接口

	go s.writeMongo()
}


type JsonStatisticsController struct {
	server.Controller
}


func (ctrl *JsonStatisticsController) noop(ctx *gin.Context)  {
	ctrl.ResponseSuccess(ctx)
}

func (ctrl *JsonStatisticsController) stat(ctx *gin.Context)  {
	numofchan := len(dataChan)
	time.Sleep( 10*time.Second )
	fmt.Println("components.E.Opts.DBConfig",components.E.Opts.DBConfig)
	ctrl.Data = fmt.Sprintf("num fo chan : %d",numofchan)
	ctrl.ResponseSuccess(ctx)
}

func (ctrl *JsonStatisticsController) c(ctx *gin.Context)  {

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Printf("exception : %s\n", r)
	//	}
	//}()

	datas := ctx.PostForm("data")
	isgz := ctx.PostForm("gzip")
	//fmt.Println("indataindataindata", xxx , eee)
	//fmt.Println("indataindataindata",datas)
	//fmt.Println("indataindataindata",isgz)

	var jsonbyte []byte
	if isgz=="1" {
		decodeBytes, err := base64.StdEncoding.DecodeString(datas)
		if err==nil{
			indata := bytes.NewBuffer(decodeBytes)
			ungz, err := gzip.NewReader(indata)
			if err==nil {
				undatas, err1 := ioutil.ReadAll(ungz)
				if err1 == nil {
					jsonbyte = undatas
				}
			}
		}
	} else {
		jsonbyte = []byte(datas)
	}


	var indata []interface {}
	err2 := json.Unmarshal(jsonbyte, &indata)

	//fmt.Println("indataindataindata",indata)
	//fmt.Println("indataindataindata",string(jsonbyte))

	if err2==nil {
		select {
		case dataChan <- indata:
			//do sth
		case <- time.After(5*time.Microsecond):
			//to sth
		}

		ctrl.Data = "ok"
		ctrl.ResponseSuccess(ctx)
	} else {
		ctrl.Data = "fail"
		ctrl.ResponseError(ctx, err2)
	}
}


func (ctrl *JsonStatisticsController) writeMongo()  {

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Printf("exception recover：%s\n", r)
	//	}
	//}()

	//mgcligetlll,_ := components.MongoGetClient()
	//mgcligetlllxx,_ := components.MongoGetClient()

	//fmt.Println(mgcligetlll,mgcligetlllxx)

	for {
		select {
			case data := <-dataChan:

				var err error
				for {
					mgcli,mgc,err = getMongoClient()
					if err==nil {
						if _, err = mgc.InsertMany(context.TODO(), data); err == nil{
							break
						}
					}
				}

			case <- time.After(3*time.Second):
				releaseMongoClient()
		}
	}
	
	fmt.Println("test writeMongo")
}

var mgc *mongo.Collection
var mgcli *components.MongoClient

func getMongoClient() (*components.MongoClient,*mongo.Collection,error)  {
	if mgcli != nil {
		return  mgcli,mgc,nil
	}
	mgcliget,err := components.MongoGetClient()
	if err==nil {
		mgcli = mgcliget
		now  := time.Now()
		collectionstr := fmt.Sprintf("%s-%d-%d-%d",collectiontableconfig,now.Year(),now.Month(),now.Day())
		mgc = mgcli.Client.Database(databaseconfig).Collection(collectionstr)
		return mgcli,mgc,nil
	}
	return nil,nil,err
}

func releaseMongoClient(){
	//fmt.Println("releaseMongoClient",mgcli)
	if mgcli != nil {
		mgcli.Release()
		mgcli = nil
	}
}