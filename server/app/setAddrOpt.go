package app

import "PortRelay/util"

//
func (s *Server) SetAddr(clientData *ClientData) {
	backDataMap := make(map[string]interface{})
	backDataMap["type"] = "set_addr_back"

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
	ResponseChan[s.Key] = make(map[string]chan []byte, 10000)
}
