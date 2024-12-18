package app

import (
	"PortRelay/util"
	"PortRelay/variable"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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
	retData := variable.ProtoHttpParam{
		ProtoCommParam: variable.ProtoCommParam{
			Proto: "http",
			UUID:  util.If(ConfigData.Server.Active == "dev", "1234567890", util.RandomUUID()),
		},
		Object: variable.HttpObjectParam{
			Host:   c.Request.Host,
			URL:    c.Request.RequestURI,
			Method: http.MethodGet,
			Header: c.Request.Header,
		},
	}

	// 处理请求
	key := util.Md5(c.Request.Host)
	jsonStr, err := json.Marshal(retData)
	if _, ok := ServerList[key]; ok && err == nil {
		ServerList[key].Write(jsonStr)
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 404,
			"msg":  "not found",
		})
	}

	//
	ResponseChan[key] = make(map[string]chan []byte)
	ResponseChan[key][retData.UUID.(string)] = make(chan []byte)
	// 读取数据
	select {
	// case clientData := <-ResponseChan[key]["retData.UUID.(string)"]:
	case clientData := <-ResponseChan[key][retData.UUID.(string)]:
		{
			c.JSON(http.StatusOK, cast.ToStringMap(string(clientData)))
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
