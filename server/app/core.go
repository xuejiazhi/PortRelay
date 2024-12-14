package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func (s *Server) Start() {
	// 解析地址
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.IPAddr, s.Port))
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
	fmt.Println("Starting PortRelay  Success...")

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
		s.conn = conn
		//读取数据
		go s.Read(conn)
	}
}

func (s *Server) Stop() {
	if s.conn == nil {
		return
	}
	s.conn.Close()
}

func (s *Server) Read(conn *net.TCPConn) {
	// 读取数据
	buffer := make([]byte, DefaultBufferSize)
	for {
		cnt, err := conn.Read(buffer)
		// 读取失败
		if err != nil {
			log.Printf("Read failed: %v", err)
			return
		}

		// 打印数据
		clientData := ClientData{}
		// 解析数据
		json.Unmarshal(buffer[:cnt], &clientData)
		// 路由
		s.Router(clientData)
	}
}

// 路由
func (s *Server) Router(clientData ClientData) {
	// 路由
	switch clientData.Type {
	case SetAddrType:
		s.SetAddr(clientData)
	}
}

func (s *Server) SetConn() {
	// TODO
}
func (s *Server) Write() {
	// TODO
}
