package main

import (
	"PortRelay/server/app"
)

func main() {
	//启动http服务
	go app.InitHttpServer()
	// select {}
	// 启动服务器
	app.NewServer("test", "127.0.0.1", 8080).Start()
}
