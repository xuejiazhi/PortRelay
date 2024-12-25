package variable

type ProtoCommParam struct {
	Proto string      `json:"proto"`
	UUID  interface{} `json:"uuid"`
}

type ProtoParam struct {
	ProtoCommParam
	Object interface{} `json:"object"`
}

type Object struct {
	Header map[string][]string `json:"header"`
	Body   []byte              `json:"body"`
}

type AddrObject struct {
	RemoteUrl string `json:"RemoteUrl"`
	LocalPort int    `json:"LocalPort"`
	LocalIP   string `json:"LocalIP"`
}

type AgentObject struct {
}

type ProtoHttpParam struct {
	ProtoCommParam
	Object HttpObjectParam `json:"object"`
}

type ProtoPostFormParam struct {
	ProtoCommParam
	Object HttpPostFormObject `json:"object"`
}

type ClientData struct {
	Type int         `json:"type"` // 类型
	Data interface{} `json:"data"` // 数据
}
