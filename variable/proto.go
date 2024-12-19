package variable

type ProtoCommParam struct {
	Proto string      `json:"proto"`
	UUID  interface{} `json:"uuid"`
}

type ProtoParam struct {
	ProtoCommParam
	Object interface{} `json:"object"`
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
	Type string      `json:"type"` // 类型
	Data interface{} `json:"data"` // 数据
}
