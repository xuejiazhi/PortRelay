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

	i := 0
	for {
		fmt.Println("client send msg:", i)
		data := `{
		    "type":"set_addr",
			"data":{
				"RemoteUrl":"https://127.0.0.1:8081",
				"LocalPort": 8100,
				"LocalIP":"127.0.0.1"
			}
		}`
		_, err := conn.Write([]byte(data))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read buf error")
			return
		}

		fmt.Printf("server call back %s,cnt= %d\n", buf, cnt)

		//cpu阻塞
		time.Sleep(1 * time.Second)
		i++
	}
}
