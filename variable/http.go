package variable

// 定义http Object传输参数
type HttpObjectParam struct {
	Host        string              `json:"host"`
	URL         string              `json:"url"`
	Method      string              `json:"method"`
	Header      map[string][]string `json:"header"`
	ContentType string              `json:"content_type"`
	Body        interface{}         `json:"body"`
}

// 定义http PostForm传输参数
type HttpPostFormObject struct {
	HttpObjectParam `json:"param"`
	ContentType     string      `json:"content_type"`
	Body            interface{} `json:"body"`
}

var (
	//	类型
	LoginType       = "login"
	LoginBackType   = "login_back"
	SetAddrType     = "set_addr"
	SetAddrBackType = "set_addr_back"
	CallBackType    = "callback"
)

const (
	ContentType_JSON                              = "application/json"
	ContentType_HTML                              = "text/html"
	ContentType_XML                               = "application/xml"
	ContentType_XML2                              = "text/xml"
	ContentType_Plain                             = "text/plain"
	ContentType_Application_X_WWW_Form_Urlencoded = "application/x-www-form-urlencoded"
	ContentType_Multipart_FormData                = "multipart/form-data"
	ContentType_PROTOBUF                          = "application/x-protobuf"
	ContentType_MSGPACK                           = "application/x-msgpack"
	ContentType_MSGPACK2                          = "application/msgpack"
	ContentType_YAML                              = "application/x-yaml"
	ContentType_YAML2                             = "application/yaml"
	ContentType_TOML                              = "application/toml"
)

var NotFound = map[string]interface{}{
	"code": 404,
	"msg":  "not found",
}

var Timeout = map[string]interface{}{
	"code": 500,
	"msg":  "timeout",
}
