package tget

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joselee214/j7f/components/http/server"
	"j7go/components"
	"j7go/models"
	"net/http"
)

func Init(g *gin.Engine) {
	s := &TgetController{}
	g.GET("/test",s.test)
}


type TgetController struct {
	server.Controller
}

func (ctrl *TgetController) test(ctx *gin.Context)  {
	param,_ := ctx.GetQuery("zzz")
	fmt.Println("=======>",param)
	dbConn, _ := components.M.GetSlaveConn()
	data,err := models.InitdbByID(dbConn,1)
	if err != nil {
		fmt.Println("======x=>",err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data, "msg": "", "result": 1})
}