package util

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"reflect"
	"unsafe"
)

func ZeroCopyByte(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Md5(str string) string {
	// 计算字符串的MD5哈希值
	hash := md5.Sum([]byte(str))
	// 将哈希值转换为十六进制字符串
	md5Str := fmt.Sprintf("%x", hash)
	// 返回MD5字符串
	return md5Str
}

// 解析url获取域名
func GetRemoteUrl(urlStr string) string {
	// 解析URL
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}

	// 获取域名
	domain := u.Hostname() + ":" + u.Port()
	return domain
}
