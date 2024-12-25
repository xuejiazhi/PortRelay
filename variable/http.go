package variable

// HttpObjectParam 定义http Object传输参数
type HttpObjectParam struct {
	Host        string              `json:"host"`
	URL         string              `json:"url"`
	Method      string              `json:"method"`
	Header      map[string][]string `json:"header"`
	ContentType string              `json:"content_type"`
	Body        interface{}         `json:"body"`
}

// HttpPostFormObject 定义http PostForm传输参数
type HttpPostFormObject struct {
	HttpObjectParam `json:"param"`
	ContentType     string      `json:"content_type"`
	Body            interface{} `json:"body"`
}

var (
	LoginType       = 0x01
	LoginBackType   = 0x02
	SetAddrType     = 0x03
	SetAddrBackType = 0x04
	CallBackType    = 0x05
)

const (
	ContentTypeJson                          = "application/json"
	ContentTypeHTML                          = "text/html"
	ContentTypeXML                           = "application/xml"
	ContentTypeXML2                          = "text/xml"
	ContentTypePlain                         = "text/plain"
	ContentTypeApplicationXWWWFormUrlencoded = "application/x-www-form-urlencoded"
	ContentTypeMultipartFormData             = "multipart/form-data"
	ContentTypePROTOBUF                      = "application/x-protobuf"
	ContentTypeMSGPACK                       = "application/x-msgpack"
	ContentTypeMSGPACK2                      = "application/msgpack"
	ContentTypeYAML                          = "application/x-yaml"
	ContentTypeYAML2                         = "application/yaml"
	ContentTypeTOML                          = "application/toml"
)

var NotFound = map[string]interface{}{
	"code": 404,
	"msg":  "not found",
}

var Timeout = map[string]interface{}{
	"code": 500,
	"msg":  "timeout",
}
