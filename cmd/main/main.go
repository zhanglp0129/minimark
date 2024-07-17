package main

import (
	"fmt"
	"minimark/config"
	"minimark/dao"
	"minimark/routers"
)

func main() {
	cfg := config.GetConfig()
	dao.GetDB()

	if cfg.Server.Https {
		runHttps()
	} else {
		runHttp()
	}
}

// 运行http
func runHttp() {
	// 获取相关配置
	cfg := config.GetConfig()
	ip := cfg.Server.IP
	port := cfg.Server.Port
	addr := fmt.Sprintf("%s:%d", ip, port)

	// 监听端口
	r := routers.GetRouters()
	err := r.Run(addr)
	if err != nil {
		panic(err)
	}
}

// 运行https
func runHttps() {
	// 获取相关配置
	cfg := config.GetConfig()
	ip := cfg.Server.IP
	port := cfg.Server.Port
	certFile := cfg.Server.CertificateFile
	keyFile := cfg.Server.CertificateKeyFile
	addr := fmt.Sprintf("%s:%d", ip, port)

	// 监听端口
	r := routers.GetRouters()
	err := r.RunTLS(addr, certFile, keyFile)
	if err != nil {
		panic(err)
	}
}
