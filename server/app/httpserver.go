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
	r.Run(fmt.Sprintf(":%d", ConfigData.Server.HttpPort))
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

	// 定义临时通道
	ResponseChan[key][retData.UUID.(string)] = make(chan interface{})
	// 读取数据
	select {
	case clientData := <-ResponseChan[key][retData.UUID.(string)]:
		{
			//删除临时通道
			delete(ResponseChan[key], retData.UUID.(string))
			cliMapData := cast.ToStringMap(clientData)
			if header, ok := cliMapData["header"]; ok {
				headerMap := cast.ToStringMap(header)
				for k, v := range headerMap {
					headValue := cast.ToStringSlice(v)
					if len(headValue) > 0 {
						c.Header(k, headValue[0])
					}
				}
			}

			// 存在body
			if body, ok := cliMapData["body"]; ok {
				c.String(http.StatusOK, cast.ToString(body))
			} else {
				//转发失败的情况
				c.String(http.StatusOK, "Port Forwarding failed!")
			}
		}
	case <-time.After(15 * time.Second):
		{
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
