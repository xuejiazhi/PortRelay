package app

import (
	"PortRelay/util"
	"PortRelay/variable"
	"encoding/json"
	"log"

	"github.com/spf13/cast"
)

func (s *Server) Login(clientData *variable.ClientData) {
	//	set backdata
	backDataMap := variable.ClientData{
		Type: variable.LoginBackType,
	}
	// 解析数据
	cliData, ok := clientData.Data.(map[string]interface{})
	if !ok {
		backDataMap.Data = variable.ErrMsg{
			ErrCode: 4000,
			ErrMsg:  "LoginData parsing failed",
		}
	} else {
		// 校验密钥secret
		secret := cast.ToString(cliData["secret"])
		decryCode, err := util.Decrypt(util.ZeroCopyByte(secret))
		if err == nil &&
			string(decryCode) == ConfigData.Server.Secret {
			backDataMap.Data = variable.ErrMsg{
				ErrCode: 200,
				ErrMsg:  "success",
			}
		} else {
			backDataMap.Data = variable.ErrMsg{
				ErrCode: 4001,
				ErrMsg:  "Secret is failed",
			}
		}
	}

	//回写
	callBackByte, err := json.Marshal(backDataMap)
	if err == nil {
		s.Conn.Write(callBackByte)
	} else {
		log.Printf("callback json is error [%v]", string(callBackByte))
	}
}
