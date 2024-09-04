package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// 配置结构体
type Config struct {
	Target    string // 目标地址
	LocalPort string // 本地端口
}

// 从命令行参数读取配置
func loadConfig() Config {
	target := flag.String("target", "https://github.com", "Proxy target address")
	localPort := flag.String("port", "80", "Local server port")
	flag.Parse()

	return Config{
		Target:    *target,
		LocalPort: *localPort,
	}
}

func main() {
	config := loadConfig()

	// 解析目标 URL
	targetUrl, err := url.Parse(config.Target)
	if err != nil {
		log.Fatalf("解析目标 URL 错误: %v", err)
	}

	// 创建反向代理
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)

	// 自定义错误处理
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		http.Error(w, "代理时发生错误", http.StatusBadGateway)
		log.Printf("代理错误: %v", err)
	}

	// 自定义修改请求
	proxy.ModifyResponse = func(resp *http.Response) error {
		// 可以在这里修改响应
		return nil
	}

	// 设置处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// log.Printf("代理请求: %s %s", r.Method, r.URL)
		r.Host = targetUrl.Host
		proxy.ServeHTTP(w, r)
	})

	// 启动服务器
	fmt.Println("反向代理服务器启动在: 127.0.0.1:80")
	if err := http.ListenAndServe(":"+config.LocalPort, nil); err != nil {
		log.Fatalf("启动服务器错误: %v", err)
	}
}
