package variable

// 定义http Object传输参数
type HttpObjectParam struct {
	Host   string              `json:"host"`
	URL    string              `json:"url"`
	Method string              `json:"method"`
	Header map[string][]string `json:"header"`
}

var (
	//	类型
	LoginType       = "login"
	LoginBackType   = "login_back"
	SetAddrType     = "set_addr"
	SetAddrBackType = "set_addr_back"
	CallBackType    = "callback"
)
