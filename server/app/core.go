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
	NewServer().Start()
}

func LoadIni() {
	// 加载ini文件
	cfg, err := ini.Load("server.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}

	// 读取配置文件
	ConfigData.Server.Name = cfg.Section("server").Key("name").MustString("test")     // 测试服务器名称
	ConfigData.Server.IP = cfg.Section("server").Key("ip").MustString("127.0.0.1")    //
	ConfigData.Server.HttpPort = cfg.Section("server").Key("http_port").MustInt(8081) // 8081
	ConfigData.Server.TcpPort = cfg.Section("server").Key("tcp_port").MustInt(8080)   // 8080
	ConfigData.Server.Active = cfg.Section("server").Key("active").MustString("dev")  // dev or prod
	ConfigData.Server.Secret = cfg.Section("server").Key("secret").MustString("12345678")
}
