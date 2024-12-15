package app

import (
	"PortRelay/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	// get params
	retData := make(map[string]interface{})
	//set retData
	retData["host"] = c.Request.Host
	retData["url"] = c.Request.RequestURI
	retData["method"] = http.MethodGet
	retData["header"] = c.Request.Header

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
	buffer := make([]byte, DefaultBufferSize)
	ServerList[util.Md5(c.Request.Host)].Read(buffer)

	// 解析数据
	clientData := make(map[string]interface{})
	bufferStr := util.GetIndexStr(strings.Trim(string(buffer), "\x00"))
	fmt.Println(bufferStr)
	err = json.Unmarshal([]byte(bufferStr), &clientData)

	// 响应数据
	c.JSON(http.StatusOK, clientData)
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
