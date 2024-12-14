package app

import "PortRelay/util"

//
func (s *Server) SetAddr(clientData ClientData) {
	// 解析数据
	cliData, ok := clientData.Data.(map[string]interface{})
	if !ok {
		s.conn.Write([]byte("SetAddrData parsing failed"))
		return
	}

	//remoteUrl
	remoteUrl, ok := cliData["RemoteUrl"]
	if !ok {
		s.conn.Write([]byte("SetAddrData parsing failed"))
		return
	}

	// 解析url
	url := util.GetRemoteUrl(remoteUrl.(string))
	if url == "" {
		s.conn.Write([]byte("SetAddrData parsing failed"))
		return
	}
	// 保存连接
	ServerList[util.Md5(url)] = s.conn
}
