package app

import (
	"log"
	"os"

	"github.com/go-ini/ini"
)

func Run() {
	//加载配置文件
	LoadIni()
	// 连接服务端
	Client{}.Dial()
}

func LoadIni() {
	// 加载ini文件
	cfg, err := ini.Load("agent.ini")
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
		return
	}

	// 读取文件内容
	content, err := os.ReadFile("agent.pem")
	if err != nil {
		log.Fatal(err)
		return
	}

	// 保存证书内容
	ConfigData.PKIXPubKey = string(content)

	// 读取配置文件
	ConfigData.Agent.Network = cfg.Section("agent").Key("network").MustString("tcp")
	ConfigData.Agent.Serverip = cfg.Section("agent").Key("serverip").MustString("127.0.0.1")
	ConfigData.Agent.Serverport = cfg.Section("agent").Key("serverport").MustInt(8080)
	ConfigData.Agent.Secret = cfg.Section("agent").Key("secret").MustString("1234567890")
	//set mapping
	ConfigData.Mapping.Name = cfg.Section("mapping").Key("name").MustString("test")
	ConfigData.Mapping.RemoteURL = cfg.Section("mapping").Key("remoteurl").MustString("http://127.0.0.1")
	ConfigData.Mapping.LocalIP = cfg.Section("mapping").Key("localip").MustString("127.0.0.1")
	ConfigData.Mapping.LocalPort = cfg.Section("mapping").Key("localport").MustInt(80)
}
