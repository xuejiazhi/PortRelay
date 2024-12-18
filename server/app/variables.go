package app

import (
	"net"
	"sync"
)

/*
*
client 发送数据格式

	{
	 "type":"set_addr"
	 "data":{
		"RemoteUrl":"https://127.0.0.1", // 远程地址
		"LocalPort": 80, // 转发本地端口
		"LocalIP":"127.0.0.1" // 转发本地地址
	 }
	}

回包数据格式

		{
		 "type":"ret_set_addr"
		 "data":{
			"code":0, // 0 成功 1 失败
			"msg":"success" // 成功或失败原因
		 }
	    }

*
*/
// type ClientData struct {
// 	Type string      `json:"type"` // 类型
// 	Data interface{} `json:"data"` // 数据
// }

type SetAddrData struct {
	RemoteUrl string // 远程地址
	LocalPort int    // 转发本地端口
	LocalIP   string // 转发本地地址
}

type LoginData struct {
	AppId     string
	AppSecret string
}

var (
	// 服务器
	ServerList        = make(map[string]*net.TCPConn)
	DefaultBufferSize = 1024 // 默认缓冲区大小

	RspLock      sync.Mutex                                           // 锁
	ResponseChan = make(map[string]map[string]chan interface{}, 1000) // 响应通道 string key and uuid
)

func NewServer() *Server {
	return &Server{
		Name:     ConfigData.Server.Name,
		IP:       ConfigData.Server.IP,
		HttpPort: ConfigData.Server.HttpPort,
		TcpPort:  ConfigData.Server.TcpPort,
	}
}

// 中继服务器
type RelayServer interface {
	Start()
	Stop()
	Read(*net.TCPConn)
}

var ConfigData Config

type Config struct {
	Server Server `ini:"server"`
}

type Server struct {
	Name     string
	IP       string
	HttpPort int
	TcpPort  int
	Active   string
	Secret   string
	Key      string
	Conn     *net.TCPConn
}
