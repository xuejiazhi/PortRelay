package app

import (
	"PortRelay/variable"
	"fmt"

	"github.com/spf13/cast"
)

func (s *Server) Callback(clientData *variable.ClientData) {
	//	解析数据
	dataMap := cast.ToStringMap(clientData.Data)
	if dataMap == nil {
		fmt.Println("Callback data parsing failed")
		return
	}

	//	解析uuid
	uuid := cast.ToString(dataMap["uuid"])
	proto := cast.ToString(dataMap["proto"])
	object := cast.ToString(dataMap["object"])
	// 将数据放入通道
	switch proto {
	case "http":
		ResponseChan[s.Key][uuid] <- []byte(object)
	}
}
