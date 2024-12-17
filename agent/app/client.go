package app

import (
	"PortRelay/util"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

type Client struct {
	Conn net.Conn `json:"conn"` // 连接
}

func (c Client) Dial() {
	// 加载配置文件
	log.Println("Agent will be Start...")
	// 建立连接
	address := fmt.Sprintf("%s:%d", ConfigData.Agent.Serverip, ConfigData.Agent.Serverport)
	log.Printf("Begin to connect to server %s\n", address)
	conn, err := net.Dial(ConfigData.Agent.Network, address)
	if err != nil {
		log.Printf("Agent start err,exit! error %v\n", err)
		return
	}

	// 连接成功
	log.Println("-> connect server success...")
	// 保存连接
	c.Conn = conn
	// 设置超时时间
	c.Conn.SetDeadline(time.Now().Add(60 * 60 * time.Second))
	// 设置读超时时间
	c.Conn.SetReadDeadline(time.Now().Add(15 * time.Second))
	// 设置写超时时间
	// c.Conn.SetWriteDeadline(time.Now().Add(15 * time.Second))

	//step1  登录
	log.Println("begin login server...")
	if err := c.Login(); err != nil {
		fmt.Printf("login fail! error %v\n", err)
		return
	}
	log.Println("-> login success!")

	//step2  设置地址
	if err = c.SetAddr(); err != nil {
		fmt.Printf("Set Address fail! error %v\n", err)
		return
	}
	log.Println("-> Set Address success!")

	go c.Read()
	select {}
}

// 登录
func (c Client) Login() error {
	// 登录
	loginData := map[string]interface{}{
		"type": "login",
		"data": map[string]interface{}{
			"secret": ConfigData.Agent.Secret,
		},
	}
	// 转换为json
	jsonByte, err := json.Marshal(loginData)
	if err != nil {
		return err
	}
	c.Conn.Write(jsonByte)

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
	typeStr, ok := buff["type"].(string)
	if !ok || typeStr != "login_back" {
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

// 设置地址
func (c Client) SetAddr() error {
	log.Printf("Start registering with remote service,Name %v", ConfigData.Mapping.Name)
	// set Address Data
	setAddrData := map[string]interface{}{
		"type": "set_addr",
		"data": map[string]interface{}{
			"RemoteUrl": ConfigData.Mapping.RemoteURL,
			"LocalPort": ConfigData.Mapping.LocalPort,
			"LocalIP":   ConfigData.Mapping.LocalIP,
		},
	}

	// 转换为json
	jsonByte, err := json.Marshal(setAddrData)
	if err != nil {
		return err
	}

	c.Conn.Write(jsonByte)

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
	typeStr, ok := buff["type"].(string)
	if !ok || typeStr != "set_addr_back" {
		return errors.New("set addr is fail")
	}

	//获取data
	data, ok := buff["data"].(map[string]interface{})
	if ok {
		errCode, ok := data["errCode"].(float64)
		if ok && errCode == 200 {
			return nil
		} else {
			//
			return fmt.Errorf("set addr is fail,errCode is %v", errCode)
		}
	} else {
		//
		return errors.New("set addr is fail")
	}
}

// Read 数据
func (c Client) Read() {
	fmt.Print(`
********************************************************
*    			开始使用HTTP隧道工具           *
********************************************************
`)
	c.Conn.SetReadDeadline(time.Now().Add(60 * 60 * time.Second))
	for {
		// 接收数据
		buf := make([]byte, 2048)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			// 连接断开
			log.Printf("read buf error %v \n ", err)
			log.Println("-> reconnect server...")
			// 重新连接服务器
			address := fmt.Sprintf("%s:%d", ConfigData.Agent.Serverip, ConfigData.Agent.Serverport)
			// 建立连接
			for {
				c.Conn, err = net.Dial(ConfigData.Agent.Network, address)
				if err != nil {
					log.Printf("reconnect server err,error %v\n", err)
					time.Sleep(5 * time.Second)
					continue
				} else {
					// 登录
					if err := c.Login(); err != nil {
						log.Printf("login fail! error %v\n", err)
						time.Sleep(5 * time.Second)
						continue
					}
					// 设置地址
					if err = c.SetAddr(); err != nil {
						log.Printf("Set Address fail! error %v\n", err)
						time.Sleep(5 * time.Second)
						continue
					}
					// 连接成功
					break
				}
			}

		}

		// 处理数据
		go c.Marshal(buf[:cnt])
	}
}

func (c Client) Marshal(buffer []byte) {
	log.Println("Recv Msg:", string(buffer))
}
