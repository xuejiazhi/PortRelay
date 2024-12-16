package main

import (
	"PortRelay/server/app"
	"PortRelay/util"
)

func main() {
	// 监听信号
	go util.SignalNotify()

	// 启动服务
	app.Run()
}
