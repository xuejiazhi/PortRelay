package app

import "encoding/json"

type Request struct {
	Header map[string][]string `json:"header"`
	Host   string              `json:"host"`
	Method string              `json:"method"`
	Url    string              `json:"url"`
	Body   string              `json:"body"`
}

func AliasRequest(reqStr string) *Request {
	// 解析请求
	var req Request
	json.Unmarshal([]byte(reqStr), &req)

	return &Request{}
}
func (r *Request) GetHost() string {
	return r.Host
}
func (r *Request) GetUrl() string {
	return r.Url
}

func (r *Request) GetMethod() string {
	return r.Method
}

func (r *Request) GetBody() string {
	return r.Body
}

func (r *Request) GetHeader() map[string][]string {
	return r.Header
}
