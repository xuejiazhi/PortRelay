package util

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

func Request(url string, method string, data string, header map[string]string) (interface{}, error) {
	httpClient := http.Client{}
	// 发起请求
	return requestDo(httpClient, url, method, data, header)
}

func RequestTLS(url string, method string, data string, header map[string]string) (interface{}, error) {
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := http.Client{Transport: tr}
	// 发起请求
	return requestDo(httpClient, url, method, data, header)
}

func requestDo(httpClient http.Client, url string, method string, data string, header map[string]string) (interface{}, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	for h, v := range header {
		req.Header.Set(h, v)
	}

	req.Header.Add("Content-Type", "applcation/json;charset=utf-8")
	response, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	} else if response != nil {
		defer response.Body.Close()

		r_body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		} else {
			return string(r_body), nil
		}
	} else {
		return nil, nil
	}
}
