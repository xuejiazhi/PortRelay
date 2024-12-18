package app

import (
	"PortRelay/util"
	"PortRelay/variable"
	"log"
)

func (s *Server) SetAddr(clientData *variable.ClientData) {
	backDataMap := make(map[string]interface{})
	backDataMap["type"] = variable.SetAddrBackType

	// 解析数据
	cliData, ok := clientData.Data.(map[string]interface{})
	if !ok {
		backDataMap["data"] = map[string]interface{}{
			"errCode": 4000,
			"errMsg":  "LoginData parsing failed",
		}
	}

	//remoteUrl
	remoteUrl, ok := cliData["RemoteUrl"]
	if !ok {
		s.Conn.Write([]byte("SetAddrData parsing failed"))
		return
	}

	// 解析url
	url := util.GetRemoteUrl(remoteUrl.(string))
	if url == "" {
		s.Conn.Write([]byte("SetAddrData parsing failed"))
		return
	}

	// 保存连接
	s.Key = util.Md5(url)
	ServerList[s.Key] = s.Conn
	ResponseChan[s.Key] = make(map[string]chan interface{})

	//success
	backDataMap["data"] = map[string]interface{}{
		"errCode": 200,
		"errMsg":  "success",
	}

	//回写
	callbackStr, _ := util.Map2Json(backDataMap)
	log.Println("Set Address Sucess Callback:", callbackStr)
	s.Conn.Write(util.ZeroCopyByte(callbackStr))

}
