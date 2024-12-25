package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cast"
)

// PostUrlEncodedForm 发送application/x-www-form-urlencoded POST请求
func PostUrlEncodedForm(posturl string, postdata map[string][]string, header map[string]interface{}) ([]byte, map[string][]string, error) {
	// 设置要发送的数据
	formData := url.Values(postdata)
	// 创建一个HTTP客户端
	client := &http.Client{}
	// 创建一个POST请求
	req, err := http.NewRequest(http.MethodPost, posturl, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, nil, err
	}

	for h, v := range header {
		req.Header.Set(h, cast.ToString(v))
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()

	// 打印响应状态码
	fmt.Println("Response Status:", resp.Status)

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, nil, err
	}

	// 打印响应体
	log.Println("Response Body:", string(body))
	// 打印响应头
	return body, resp.Header, nil
}

// PostMultiForm 发送multipart/form-data POST请求
func PostMultiForm(posturl string, postdata, header map[string]interface{}) ([]byte, map[string][]string, error) {
	// 设置要发送的数据
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if postdata != nil || len(postdata) > 0 {
		for k, v := range postdata {
			_ = writer.WriteField(k, cast.ToString(v))
		}
	}

	// 关闭multipart.Writer
	_ = writer.Close()
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建一个POST请求
	req, err := http.NewRequest("POST", posturl, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, nil, err
	}

	// 设置请求头的Content-Type为multipart/form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()

	// 打印响应状态码
	fmt.Println("Response Status:", resp.Status)

	// 读取响应体
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, nil, err
	}

	// 打印响应体
	fmt.Println("Response Body:", string(bodyBytes))

	// 打印响应头
	return bodyBytes, resp.Header, nil
}

// 发送application/json POST请求
func PostJson(posturl string, postdata, header map[string]interface{}) ([]byte, map[string][]string, error) {
	// 设置要发送的数据
	jsonBytes, _ := json.Marshal(postdata)
	// 创建一个HTTP客户端
	client := &http.Client{}
	// 创建一个POST请求
	req, err := http.NewRequest(http.MethodPost, posturl, bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, nil, err
	}
	for h, v := range header {
		req.Header.Set(h, cast.ToString(v))
	}

	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()
	// 打印响应状态码
	fmt.Println("Response Status:", resp.Status)
	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, nil, err
	}
	// 打印响应体
	log.Println("Response Body:", string(body))
	// 打印响应头
	return body, resp.Header, nil
}
