package main

import (
	"NewBlock/handler"
	"NewBlock/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	log "github.com/sirupsen/logrus"
	"myCommon/common/middleware"
	"time"
)

func main(){
	service:=web.NewService(
		web.Name("block"),
		web.Version("latest"),
		web.Address(":12345"),
	)

	service.Init()

	t:=time.Now()
	genesisBlock:=models.Block{}
	genesisBlock=models.Block{
		Index:     0,
		Timestamp: t.String(),
		BPM:       0,
		Hash:      models.CalculateHash(genesisBlock),
		PrevHash:  "",
	}
	spew.Dump(genesisBlock)

	models.BlockChain=append(models.BlockChain,genesisBlock)

	g:=new(handler.Gin)
	gin.SetMode(gin.ReleaseMode)
	router:=gin.Default()
	router.Use(middleware.Cors())

	router.GET("/",g.GetBlockChain)
	router.POST("/",g.WriteBlock)

	// 注册
	service.Handle("/",router)

	// 运行
	err:=service.Run()
	if err!=nil{
		log.Error("服务启动失败: ",err)
		return
	}
}