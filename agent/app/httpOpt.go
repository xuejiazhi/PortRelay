package app

import (
	"PortRelay/util"
	"PortRelay/variable"
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

func (h *HttpOpt) Analysis() ([]byte, map[string][]string, error) {
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
			return h.Post()
		case http.MethodPut:
			// h.Put()
		case http.MethodDelete:
			// h.Delete()
		default:
			// h.Option()
		}
	}

	// 错误
	return nil, nil, errors.New("http method is not support")
}

func (h *HttpOpt) Get() ([]byte, map[string][]string, error) {
	// 拼接url
	u := HostRouterList[util.Md5(h.Host)]
	var url string
	// 拼接url
	if u.Port == 80 && h.Url == "/" {
		url = fmt.Sprintf("http://%s", u.Host)
	} else {
		url = fmt.Sprintf("http://%s:%d%s", u.Host, u.Port, h.Url)
	}

	// 打印url
	log.Printf("Do Get url is %s", url)
	// 发送请求
	body, header, err := util.Request(url, http.MethodGet, "", func() map[string]string {
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
	//将数据进行压缩
	bodyByte, _ := util.CompressString(body)

	// 返回数据
	return bodyByte, header, err
}

func (h *HttpOpt) Post() ([]byte, map[string][]string, error) {
	// 拼接url
	u := HostRouterList[util.Md5(h.Host)]
	// 拼接url
	postUrl := fmt.Sprintf("http://%s:%d%s", u.Host, u.Port, h.Url)
	// 打印url
	log.Printf("Do Post url is %s", postUrl)

	// 取header
	header := func() map[string]interface{} {
		var x = make(map[string]interface{})
		if len(h.Header) > 0 {
			for k, v := range h.Header {
				if len(v) > 0 {
					x[k] = v[0]
				}
			}
		}
		return x
	}()

	//取Object的content_type
	contentType, ok := h.Object.(map[string]interface{})["content_type"].(string)
	if ok {
		switch contentType {
		case variable.ContentTypeMultipartFormData:
			{
				// 解析body
				body := cast.ToStringMap(cast.ToStringMap(h.Object)["body"])
				return util.PostMultiForm(postUrl, body, header)
			}
		case variable.ContentTypeApplicationXWWWFormUrlencoded:
			{
				// 解析body
				body := cast.ToStringMapStringSlice(cast.ToStringMap(h.Object)["body"])
				return util.PostUrlEncodedForm(postUrl, body, header)
			}
		case variable.ContentTypeJson:
			fallthrough
		default:
			{
				//获取body
				body := cast.ToStringMap(cast.ToStringMap(h.Object)["body"])
				return util.PostJson(postUrl, body, header)
			}

		}
	}
	return nil, nil, errors.New("content_type is not support")
}

func (h *HttpOpt) Put() {

}

func (h *HttpOpt) Delete() {

}

func (h *HttpOpt) Option() {

}
