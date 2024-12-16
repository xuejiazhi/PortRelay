package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client start err,exit!")
		return
	}
	data := `{
		"type":"set_addr",
		"data":{
			"RemoteUrl":"https://127.0.0.1:8081",
			"LocalPort": 8100,
			"LocalIP":"127.0.0.1"
		}
	}`
	_, err = conn.Write([]byte(data))
	if err != nil {
		fmt.Println("write conn err", err)
		return
	}

	i := 0
	for {
		// 接收数据
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read buf error")
			return
		}

		//

		conn.Write([]byte(`
		<<begin>>{
			"code":200,
			"msg":"ok",
			"time":"` + time.Now().GoString() + `"
		}<<end>>`))
		fmt.Printf("server call back %s,cnt= %d\n", buf, cnt)

		//cpu阻塞
		time.Sleep(1 * time.Second)
		i++
	}
}
