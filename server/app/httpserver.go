package app

import (
	"PortRelay/util"
	"PortRelay/variable"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
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
			DoPost(c)
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
	retData := getCommRetData(c)

	// 处理请求
	key := util.Md5(c.Request.Host)
	jsonStr, err := json.Marshal(retData)
	if _, ok := ServerList[key]; ok && err == nil {
		ServerList[key].Write(jsonStr)
	} else {
		c.JSON(http.StatusOK, variable.NotFound)
	}

	// 读取数据
	rspReadData(c, key, retData.UUID.(string))
}

func DoPost(c *gin.Context) {
	// set retData
	retData := getCommRetData(c)

	//判断是否是Header的ContentType
	if formData, err := getFormData(c); err == nil {
		retData.Object.Body = formData
	} else {
		log.Printf("getFormData err: %v \n", err)
	}

	// 处理请求
	key := util.Md5(c.Request.Host)
	if err := callAgent(key, retData); err != nil {
		c.JSON(http.StatusOK, variable.NotFound)
	}

	// 读取数据
	rspReadData(c, key, retData.UUID.(string))
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

// =============================================================================
// 处理不同的 Content-Type 类型的 POST 请求
func getFormData(c *gin.Context) (interface{}, error) {
	// 处理不同的 Content-Type 类型的 POST 请求
	switch c.ContentType() {
	case gin.MIMEPOSTForm:
		{
			// 手动解析请求体
			if err := c.Request.ParseForm(); err != nil {
				return nil, err
			}
			// 处理 application/x-www-form-urlencoded 类型的 POST 请求
			formData := c.Request.PostForm
			return formData, nil
		}
	case gin.MIMEJSON:
		{
			// 处理 application/json 类型的 POST 请求
			var jsonData map[string]interface{}
			if err := c.ShouldBindJSON(&jsonData); err != nil {
				fmt.Println("Error binding JSON:", err)
				return nil, err
			}
			return jsonData, nil
		}
	case gin.MIMEXML:
		{
			// 处理 application/xml 类型的 POST 请求
			var xmlData map[string]interface{}
			// 解析 XML 数据
			if err := c.ShouldBindXML(&xmlData); err != nil {
				fmt.Println("Error binding XML:", err)
				return nil, err
			}
			return xmlData, nil
		}
	case gin.MIMEYAML:
		{
			// 处理 application/x-yaml 类型的 POST 请求
			var yamlData map[string]interface{}
			if err := c.ShouldBindYAML(&yamlData); err != nil {
				fmt.Println("Error binding YAML:", err)
				return nil, err
			}
			return yamlData, nil
		}
	case gin.MIMEHTML:
		{
			// 处理 text/html 类型的 POST 请求
			var htmlData string
			if err := c.ShouldBind(&htmlData); err != nil {
				fmt.Println("Error binding HTML:", err)
				return nil, err
			}
			return htmlData, nil
		}
	case gin.MIMEMultipartPOSTForm:
		{
			// 处理 multipart/form-data 类型的 POST 请求
			err := c.Request.ParseMultipartForm(32 << 20) // 32MB 是默认的内存限制
			if err != nil {
				fmt.Println("Error parsing form-data:", err)
				return nil, err
			}
			// 获取所有 form-data 数据
			formData := c.Request.MultipartForm.Value
			return formData, nil
		}
	default:
		return nil, nil
	}
}

// 处理请求
func callAgent(key string, retData interface{}) error {
	// 处理请求
	jsonStr, err := json.Marshal(retData)
	// 处理请求
	if _, ok := ServerList[key]; ok && err == nil {
		ServerList[key].Write(jsonStr)
	}
	return err
}

// 通用返回数据
func getCommRetData(c *gin.Context) variable.ProtoHttpParam {
	// set retData
	return variable.ProtoHttpParam{
		ProtoCommParam: variable.ProtoCommParam{
			Proto: "http",
			UUID:  util.If(ConfigData.Server.Active == "dev", "1234567890", util.RandomUUID()),
		},
		Object: variable.HttpObjectParam{
			Host:        c.Request.Host,
			URL:         c.Request.RequestURI,
			Method:      c.Request.Method,
			Header:      c.Request.Header,
			ContentType: c.ContentType(),
		},
	}
}
func rspReadData(c *gin.Context, key, uuid string) {
	if _, ok := ResponseChan[key]; !ok {
		return
	}
	// 定义临时通道
	ResponseChan[key][uuid] = make(chan interface{})
	// 读取数据
	select {
	case clientData := <-ResponseChan[key][uuid]:
		{
			//删除临时通道
			delete(ResponseChan[key], uuid)
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
				decoded, _ := base64.StdEncoding.DecodeString(body.(string))
				c.String(http.StatusOK, string(decoded))
			} else {
				//转发失败的情况
				c.String(http.StatusOK, "Port Forwarding failed!")
			}
		}
	case <-time.After(15 * time.Second):
		{
			c.JSON(http.StatusOK, variable.Timeout)
		}
	}
}
