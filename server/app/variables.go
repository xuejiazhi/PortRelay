package app

import "net"

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
type ClientData struct {
	Type string      `json:"type"` // 类型
	Data interface{} `json:"data"` // 数据
}

type SetAddrData struct {
	RemoteUrl string // 远程地址
	LocalPort int    // 转发本地端口
	LocalIP   string // 转发本地地址
}

var (
	// 服务器
	ServerList        = make(map[string]*net.TCPConn)
	DefaultBufferSize = 1024 // 默认缓冲区大小

	//
	SetAddrType = "set_addr"
)

func NewServer(name, ipAddr string, port int) *Server {
	return &Server{
		Name:   name,
		IPAddr: ipAddr,
		Port:   port,
	}
}

// 中继服务器
type RelayServer interface {
	Start()
	Read(*net.TCPConn)
}

type Server struct {
	Name   string       // 名称
	IPAddr string       // 地址
	Port   int          // 端口
	conn   *net.TCPConn // 连接
}
