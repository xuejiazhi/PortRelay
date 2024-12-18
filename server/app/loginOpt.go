package app

import (
	"PortRelay/util"
	"PortRelay/variable"
)

func (s *Server) Login(clientData *variable.ClientData) {
	backDataMap := make(map[string]interface{})
	backDataMap["type"] = variable.LoginBackType
	// 解析数据
	cliData, ok := clientData.Data.(map[string]interface{})
	if !ok {
		backDataMap["data"] = map[string]interface{}{
			"errCode": 4000,
			"errMsg":  "LoginData parsing failed",
		}
	} else {
		// 校验密钥secret
		secret, ok := cliData["secret"].(string)
		if !ok {
			backDataMap["data"] = map[string]interface{}{
				"errCode": 4001,
				"errMsg":  "Secret is failed",
			}
		} else {
			decryCode, err := util.Decrypt(util.ZeroCopyByte(secret))
			if err == nil &&
				string(decryCode) == ConfigData.Server.Secret {
				backDataMap["data"] = map[string]interface{}{
					"errCode": 200,
					"errMsg":  "Success",
				}
			} else {
				backDataMap["data"] = map[string]interface{}{
					"errCode": 4001,
					"errMsg":  "Secret is failed",
				}
			}
		}
	}

	//回写
	callbackStr, _ := util.Map2Json(backDataMap)
	s.Conn.Write(util.ZeroCopyByte(callbackStr))
}
