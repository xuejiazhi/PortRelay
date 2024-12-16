package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"github.com/google/uuid"
)

// PprofListen ---------------------------
//
// pprof 监听
func PprofListen(port int) {
	go func() {
		_ = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
}

// SignalNotify get signal
func SignalNotify() {
	sigChan := make(chan os.Signal, 1)
	//获取信号
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	// 阻塞等待信号
	select {
	case s := <-sigChan:
		log.Fatal("os is exit,signal is:", s)
		os.Exit(0)
	}
}

func If(condition bool, x, y any) any {
	if condition {
		return x
	}
	return y
}

func MyTry(f func(), hErr func(e_ interface{})) {
	defer func() {
		if e := recover(); e != nil {
			hErr(e)
		}
	}()

	f()
}

// RandomUUID generates and returns a new random UUID as a string.
func RandomUUID() string {
	// 生成一个随机的UUID
	randomUUID := uuid.New()
	return randomUUID.String()
}

func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover: panic info:", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func Json2Map(jsonStr string) (map[string]interface{}, error) {
	// 定义一个空的map来存储解码后的数据
	var dataMap map[string]interface{}

	// 使用json.Unmarshal将JSON字符串解码到map中
	err := json.Unmarshal([]byte(jsonStr), &dataMap)
	if err != nil {
		fmt.Println("解析错误:", err)
	}
	return dataMap, err
}

func Map2Json(dataMap map[string]interface{}) (string, error) {

	// 使用json.Unmarshal将JSON字符串解码到map中
	jsonByte, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println("解析错误:", err)
	}
	return string(jsonByte), err
}

// 校验map中是否存在指定key
func JudgeMap(dataMap map[string]interface{}, keys []string) bool {
	for _, key := range keys {
		if _, ok := dataMap[key]; !ok {
			return false
		}
	}
	return true
}

// Recover is a function designed to handle panics that occur within the goroutine where it is called.
// It uses the defer keyword to ensure that the anonymous function is executed after the function completes,
// regardless of whether it completes normally or panics.
func Recover() {
	if r := recover(); r != nil {
		log.Println("Recovered from panic in goroutine:", r)
	}
}
