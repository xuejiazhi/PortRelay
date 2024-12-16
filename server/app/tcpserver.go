package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func (s *Server) Start() {
	// 解析地址
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.IP, s.TcpPort))
	if err != nil {
		log.Printf("Address parsing failed: %v", err)
		return
	}
	// 监听
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Printf("Monitoring failed: %v", err)
		return
	}

	// 启动成功
	log.Printf("Starting %v  Success...", s.Name)

	// 停止`Stop`
	defer s.Stop()
	// 开始监听
	for {
		// 接收连接
		conn, err := listen.AcceptTCP()
		// 接收失败
		if err != nil {
			log.Printf("Accept failed: %v", err)
			return
		}
		// 保存连接
		s.Conn = conn
		//读取数据
		go s.Read()
	}
}

func (s *Server) Stop() {
	// 关闭连接加锁
	RspLock.Lock()
	defer RspLock.Unlock()

	// 从服务器列表中删除
	delete(ServerList, s.Key)
	delete(ResponseChan, s.Key)

	// 关闭连接
	if s.Conn == nil {
		return
	}
	s.Conn.Close()
}

func (s *Server) Read() {
	// 读取数据
	buffer := make([]byte, DefaultBufferSize)
	for {
		cnt, err := s.Conn.Read(buffer)
		// 读取失败
		if err != nil {
			log.Printf("Read failed: %v", err)
			return
		}
		// 路由 (todo: 优化使用协程池)
		go s.Router(buffer[:cnt])
	}
}

// 路由
func (s *Server) Router(buffer []byte) {
	// 打印数据
	clientData := ClientData{}
	// 解析数据
	json.Unmarshal(buffer, &clientData)
	// 路由
	switch clientData.Type {
	case SetAddrType:
		s.SetAddr(&clientData)
	case LoginType:
		s.Login(&clientData)
	default:
		fmt.Println("unknown type", string(buffer))
	}
}

func (s *Server) SetConn() {
	// TODO
}
func (s *Server) Write() {
	// TODO
}
