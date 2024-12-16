package app

import (
	"PortRelay/util"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitHttpServer() {
	//设置gin模式
	gin.SetMode(gin.ReleaseMode)
	//创建路由
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet: //GET请求
			DoGet(c)
		case http.MethodPost: //POST请求
			DoPost()
		case http.MethodDelete: //DELETE请求
			DoDelete()
		}
	})

	fmt.Println("http server start")

	//监听端口，默认绑定端口8081
	r.Run(":8081")
}

func DoGet(c *gin.Context) {
	// set retData
	retData := map[string]interface{}{
		"type": "http",
		"uuid": util.If(ConfigData.Server.Active == "dev", "1234567890", util.RandomUUID()),
		"object": map[string]interface{}{
			"host":   c.Request.Host,
			"url":    c.Request.RequestURI,
			"method": http.MethodGet,
			"header": c.Request.Header,
		},
	}

	// 处理请求
	jsonStr, err := json.Marshal(retData)
	if _, ok := ServerList[util.Md5(c.Request.Host)]; ok && err == nil {
		ServerList[util.Md5(c.Request.Host)].Write(jsonStr)
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 404,
			"msg":  "not found",
		})
	}

	// 读取数据
	select {
	case clientData := <-ResponseChan[util.Md5(c.Request.Host)][retData["uuid"].(string)]:
		{
			fmt.Println(clientData)
			c.JSON(http.StatusOK, clientData)
		}
	case <-time.After(15 * time.Second):
		{
			fmt.Println("timeout")
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 500,
				"msg":  "timeout",
			})
		}
	}
}

func DoPost() {
	//todo
	fmt.Println("this is Post Method")
}

func DoDelete() {
	//todo
}

func DoPut() {
	//todo
}

func DoPatch() {
	//todo
}

func DoOption() {
	//todo
}
