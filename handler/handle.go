package handler

import (
	"NewBlock/models"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Gin struct {
}

func (g *Gin) GetBlockChain(c *gin.Context) {
	// marshalIndent 方法是对JSON数据进行一些格式化的处理
	body, err := json.MarshalIndent(models.BlockChain, "", " ")
	if err != nil {
		log.Error("数据接收失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    "",
			"message": err,
		})
		return
	}

	res := []interface{}{}
	json.Unmarshal(body, &res)

	log.Info("success!!!")
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    res,
		"message": nil,
	})
}

func (g *Gin) WriteBlock(c *gin.Context) {
	var req models.PostReq
	var preBlock = models.BlockChain[len(models.BlockChain)-1]
	body,err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("获取数据失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"data":    "",
			"message": err,
		})
		return
	}
	json.Unmarshal(body,&req)

	block, err := models.GenerateBlock(preBlock, req.BPM)
	if err != nil {
		log.Error("创建块失败: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    "",
			"message": err,
		})
		return
	}

	if models.IsBlockValid(block, preBlock) {
		models.BlockChain = append(models.BlockChain, block)
		// 使用spew.Dump 这个函数可以以非常美观和方便阅读的方式将 struct、slice 等数据打印在控制台里
		spew.Dump(models.BlockChain)
		//c.JSON(http.StatusOK, gin.H{
		//	"code":    http.StatusCreated,
		//	"data":    block,
		//	"message": nil,
		//})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    block,
		"message": nil,
	})
}
