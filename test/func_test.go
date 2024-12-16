package test

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func Test_Url(t *testing.T) {
	// 示例URL
	urlStr := "https://39.101.79.205:9999"

	// 解析URL
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// 获取域名
	domain := u.Hostname() + ":" + u.Port()
	fmt.Println("Domain:", domain)
}

func Test_md5(t *testing.T) {
	str := "hello world"
	hash := md5.Sum([]byte(str))

	// 将哈希值转换为十六进制字符串
	md5Str := fmt.Sprintf("%x", hash)
	fmt.Println(md5Str)
}

func TestXxx(t *testing.T) {
	str := `	{
		"type":"set_addr",
		"data":{
			"RemoteUrl":"https://39.101.79.205:9999",
			"LocalPort": 8100,
			"LocalIP":"127.0.0.1"
		}
	}`
	fmt.Println(str)
	type ClientData struct {
		Type string      `json:"type"` // 类型
		Data interface{} `json:"data"` // 数据
	}
	clientData := ClientData{}
	json.Unmarshal([]byte(str), &clientData)
	fmt.Println(clientData)
}

func Test_Xxx(t *testing.T) {
	// 假设有一个字符串
	str := `<<begin>>
		{
			"code":200,
			"msg":"ok",
			"time":"2022-01-01 12:00:00"
		}<<end>><<begin>>
		{
			"code":200,
			"msg":"ok",
			"time":"2022-01-01 12:00:00"
		}<<end>><<begin>>
		{
			"code":200,
			"msg":"ok",
			"time":"2022-01-01 12:00:00"
		}<<end>>`

	// 找到 [[begin]] 和 [[end]] 的位置
	beginIndex := strings.Index(str, "<<begin>>")
	endIndex := strings.Index(str, "<<end>>")

	// 如果找到了 [[begin]] 和 [[end]]
	if beginIndex != -1 && endIndex != -1 {
		// 提取中间的值
		middleValue := str[beginIndex+len("[[begin]]") : endIndex]
		fmt.Println("Middle Value:", middleValue)
	} else {
		fmt.Println("Begin or end marker not found")
	}
}
