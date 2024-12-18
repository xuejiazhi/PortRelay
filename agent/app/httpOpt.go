package app

import (
	"PortRelay/util"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cast"
)

type HttpOpt struct {
	Object interface{}         `json:"object"`
	Host   string              `json:"host"`
	Url    string              `json:"url"`
	Header map[string][]string `json:"header"`
}

func (h *HttpOpt) Analysis() (interface{}, error) {
	if h.Object != nil {
		// 解析数据
		objectMap := cast.ToStringMap(h.Object)
		// set host
		if host, ok := objectMap["host"]; ok {
			h.Host = cast.ToString(host)
		}

		// set url
		if url, okurl := objectMap["url"]; okurl {
			h.Url = cast.ToString(url)
		}

		// set header
		if header, ok := objectMap["header"]; ok {
			h.Header = cast.ToStringMapStringSlice(header)
		}

		//judge method
		switch objectMap["method"] {
		case http.MethodGet:
			return h.Get()
		case http.MethodPost:
			// h.Post()
		case http.MethodPut:
			// h.Put()
		case http.MethodDelete:
			// h.Delete()
		default:
			// h.Option()
		}
	}

	// 错误
	return nil, errors.New("http method is not support")
}

func (h *HttpOpt) Get() (interface{}, error) {
	// 拼接url
	u := HostRouterList[util.Md5(h.Host)]
	// 拼接url
	url := fmt.Sprintf("http://%s:%d%s", u.Host, u.Port, h.Url)
	// 打印url
	log.Printf("Do Get url is %s", url)
	// 发送请求
	return util.Request(url, http.MethodGet, "", func() map[string]string {
		var x = make(map[string]string)
		if len(h.Header) > 0 {
			for k, v := range h.Header {
				if len(v) > 0 {
					x[k] = v[0]
				}
			}
		}
		return x
	}())
}

func (h *HttpOpt) Post() {

}

func (h *HttpOpt) Put() {

}

func (h *HttpOpt) Delete() {

}

func (h *HttpOpt) Option() {

}
