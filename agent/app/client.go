package app

import (
	"PortRelay/util"
	"PortRelay/variable"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/spf13/cast"
)

type Client struct {
	Conn net.Conn `json:"conn"` // 连接
}

func (c Client) Dial() {
	// 加载配置文件
	log.Println("Agent will be Start...")

	// 建立连接
	address := fmt.Sprintf("%s:%d", ConfigData.Agent.Serverip, ConfigData.Agent.Serverport)
	conn, err := net.Dial(ConfigData.Agent.Network, address)
	if err != nil {
		log.Printf("Agent start err,exit! error %v\n", err)
		return
	}

	// 保存连接
	c.Conn = conn
	// 设置超时时间
	_ = c.Conn.SetDeadline(time.Now().Add(60 * 60 * time.Second))
	// 设置读超时时间
	_ = c.Conn.SetReadDeadline(time.Now().Add(15 * time.Second))

	//step1  登录
	if err := c.Login(); err != nil {
		fmt.Printf("-> login fail! error %v\n", err)
		return
	}
	log.Println("-> login success!")

	//step2  设置地址
	if err = c.SetAddr(); err != nil {
		fmt.Printf("Set Address fail! error %v\n", err)
		return
	}
	log.Println("-> Set Address success!")

	// 读取数据
	go c.Read()
	// 阻塞
	select {}
}

// Login 登录
func (c Client) Login() error {
	// 登录
	loginData := variable.ClientData{
		Type: variable.LoginType,
		Data: map[string]interface{}{
			"secret": ConfigData.Agent.Secret,
		},
	}
	// 转换为json
	jsonByte, err := json.Marshal(loginData)
	if err != nil {
		return err
	}

	//print log
	log.Printf("Set Login String:[%v]\n", string(jsonByte))
	//压缩
	compressByte, _ := util.CompressString(string(jsonByte))
	//编码
	callBackByte, _ := util.Encode(string(compressByte))
	_, _ = c.Conn.Write(callBackByte)

	// 读取数据
	buf := make([]byte, 1024)
	n, err := c.Conn.Read(buf)
	if err != nil {
		return err
	}

	// 解析数据
	buff, err := util.Json2Map(string(buf[:n]))
	if err != nil {
		return err
	}

	// 判断type
	loginBackType := cast.ToInt(buff["type"])
	if loginBackType != variable.LoginBackType {
		return errors.New("login is fail")
	}

	//获取data
	data, ok := buff["data"].(map[string]interface{})
	if ok {
		errCode, ok := data["errCode"].(float64)
		if ok && errCode == 200 {
			return nil
		} else {
			//
			return fmt.Errorf("login is fail,errCode is %v", errCode)
		}
	} else {
		//
		return errors.New("login is fail")
	}
}

// SetAddr 设置地址
func (c Client) SetAddr() error {
	// set Address Data
	setAddrData := variable.ClientData{
		Type: variable.SetAddrType,
		Data: variable.AddrObject{
			RemoteUrl: ConfigData.Mapping.RemoteURL,
			LocalPort: ConfigData.Mapping.LocalPort,
			LocalIP:   ConfigData.Mapping.LocalIP,
		},
	}

	// 转换为json
	jsonByte, err := json.Marshal(setAddrData)
	if err != nil {
		return err
	}

	//压缩
	compressByte, _ := util.CompressString(string(jsonByte))
	log.Printf("Set Address String:[%v]\n", string(jsonByte))
	//将消息编码
	callBackByte, _ := util.Encode(string(compressByte))
	_, _ = c.Conn.Write(callBackByte)

	//	发送数据
	_, _ = c.Conn.Write(jsonByte)

	// 读取数据
	buf := make([]byte, 1024)
	n, err := c.Conn.Read(buf)
	if err != nil {
		return err
	}

	// 解析数据
	buff, err := util.Json2Map(string(buf[:n]))
	if err != nil {
		return err
	}

	// 判断type
	addrType := cast.ToInt(buff["type"])
	if addrType != variable.SetAddrBackType {
		return errors.New("set addr is fail")
	}

	//获取data
	data := cast.ToStringMap(buff["data"])
	errCode, ok := data["errCode"].(float64)
	if ok && errCode == 200 {
		// 设置路由列表
		key := util.Md5(util.GetRemoteUrl(ConfigData.Mapping.RemoteURL))
		HostRouterList[key] = HttpRouter{
			Host: ConfigData.Mapping.LocalIP,
			Port: ConfigData.Mapping.LocalPort,
		}
		// 成功
		return nil
	} else {
		//
		return fmt.Errorf("set addr is fail,errCode is %v", errCode)
	}
}

// Read 数据
func (c Client) Read() {
	// 打印
	fmt.Print(`
********************************************************
*    		Start using HTTP tunneling tools           *
********************************************************
`)
	// 设置读取超时时间
	_ = c.Conn.SetReadDeadline(time.Now().Add(60 * 60 * time.Second))
	for {
		// 接收数据
		buf := make([]byte, 2048)
		cnt, err := c.Conn.Read(buf)
		// 处理错误 重连

		if err != nil {
			// 连接断开
			log.Printf("Agent is DisConnect,Error Message %v \n ", err)
			// 重新连接服务器
			address := fmt.Sprintf("%s:%d", ConfigData.Agent.Serverip, ConfigData.Agent.Serverport)
			for {
				// 建立连接
				if c.Conn, err = net.Dial(ConfigData.Agent.Network, address); err != nil {
					log.Printf("ReConnect Server Error %v\n", err)
					time.Sleep(3 * time.Second)
					continue
				}

				// Reconnect
				if err := c.Reconnect(); err != nil {
					time.Sleep(3 * time.Second)
					continue
				}
				// 连接成功
				log.Println("-> Connect Server Success...")
				// 退出循环
				break
			}
		}

		// 处理数据(优化使用协程池)
		go c.Marshal(buf[:cnt])
	}
}

func (c Client) Marshal(buffer []byte) {
	// 打印数据
	log.Println("Recv Msg:", string(buffer))

	// 解析数据
	bufData := variable.ProtoParam{}
	err := json.Unmarshal(buffer, &bufData)
	// 处理错误
	if err != nil {
		log.Printf("Unmarshal error->%v", err.Error())
		return
	}

	// 处理数据
	pro := ProtoTransfer(bufData.Proto, bufData.Object)
	if pro != nil {
		// 解析数据
		rspBody, rspHeader, err := pro.Analysis()
		if err != nil {
			log.Println("Analysis error:", err)
		} else {
			// 回调数据
			callback := variable.ClientData{
				Type: variable.CallBackType,
				Data: variable.ProtoParam{
					Object: variable.Object{
						Header: rspHeader,
						Body:   rspBody,
					},
					ProtoCommParam: variable.ProtoCommParam{
						Proto: bufData.Proto,
						UUID:  bufData.UUID,
					},
				},
			}

			// 转换为json
			if callbackData, err := json.Marshal(callback); err == nil {
				//压缩字符串
				compressByte, _ := util.CompressString(string(callbackData))
				//消息编码
				callBackByte, _ := util.Encode(string(compressByte))
				fmt.Println("CallBack Encode String=>", string(callBackByte))
				// 发送数据
				_, _ = c.Conn.Write(callBackByte)
			}
		}
	}
}

// Reconnect 重连
func (c Client) Reconnect() error {
	//定义错误
	var err error

	// 登录
	if err := c.Login(); err != nil {
		log.Printf("login fail! error %v\n", err)
		return err
	}

	// 设置地址
	if err = c.SetAddr(); err != nil {
		log.Printf("Set Address fail! error %v\n", err)
		return err
	}

	//return
	return nil
}
