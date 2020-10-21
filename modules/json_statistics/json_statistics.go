package json_statistics



import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joselee214/j7f/components/http/server"
	"io/ioutil"
	"compress/gzip"
	"net/http"
	"time"
)

var dataChan chan []interface {}
var chanlen int = 2000000

func Init(g *gin.Engine) {

	dataChan = make(chan []interface {},chanlen)

	s := &JsonStatisticsController{}
	g.GET("/favicon.ico",s.noop )//注册接口
	g.Any("/c",s.c)
	g.Any("/stat",s.stat )//注册接口

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
	ctrl.Data = fmt.Sprintf("num fo chan : %d",numofchan)
	ctrl.ResponseSuccess(ctx)
}

func (ctrl *JsonStatisticsController) c(ctx *gin.Context)  {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("exception : %s\n", r)
		}
	}()

	datas := ctx.PostForm("data")
	isgz := ctx.PostForm("gzip")

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

	if err2==nil {
		select {
		case dataChan <- indata:
			//do sth
		case <- time.After(5*time.Microsecond):
			//to sth
		}
	}


	ctrl.Data = "ok"
	ctrl.ResponseSuccess(ctx)

}


func (ctrl *JsonStatisticsController) writeMongo()  {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("exception recover：%s\n", r)
		}
	}()



}