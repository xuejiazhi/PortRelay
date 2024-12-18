package main

import (
	"PortRelay/agent/app"
	"PortRelay/util"
)

func main() {
	// 监听信号
	go util.SignalNotify()

	// 启动
	app.Run()
}
