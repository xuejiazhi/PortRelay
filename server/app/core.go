package app

import (
	"fmt"

	"github.com/go-ini/ini"
)

// 运行
func Run() {
	//加载配置文件
	LoadIni()

	//启动http服务
	go InitHttpServer()

	// 启动服务器
	NewServer(ConfigData.Service.Name, ConfigData.Service.IP, ConfigData.Service.TcpPort).Start()
}

func LoadIni() {
	// 加载ini文件
	cfg, err := ini.Load("server.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}

	// 读取配置文件
	ConfigData.Service.Name = cfg.Section("service").Key("name").MustString("test")
	ConfigData.Service.IP = cfg.Section("service").Key("ip").MustString("127.0.0.1")
	ConfigData.Service.HttpPort = cfg.Section("service").Key("http_port").MustInt(8081)
	ConfigData.Service.TcpPort = cfg.Section("service").Key("tcp_port").MustInt(8080)
}
