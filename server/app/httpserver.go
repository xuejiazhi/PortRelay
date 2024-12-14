package app

import (
	"PortRelay/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHttpServer() {
	//将应用切换到“发布模式”以提升性能
	gin.SetMode(gin.ReleaseMode)
	//创建路由
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet: //GET请求
			DoGet(c)
		}
	})

	fmt.Println("http server start")

	//监听端口，默认绑定端口8081
	r.Run(":8081")
}

func DoGet(c *gin.Context) {
	//
	host := c.Request.Host
	url := c.Request.RequestURI
	fmt.Println("this is Get Method:", url)
	retData := make(map[string]interface{})
	retData["host"] = host
	retData["method"] = http.MethodGet
	if len(c.Request.Header) > 0 {
		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}
	}
	c1 := ServerList
	fmt.Print(c1)
	if _, ok := ServerList[util.Md5(host)]; ok {

		ServerList[util.Md5(host)].Write(util.ZeroCopyByte("hello world"))
	}
	c.JSON(http.StatusOK, gin.H{"user": "123", "secret": "456"})
}

func DoPost() {
	fmt.Println("this is Post Method")
}
