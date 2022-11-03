package main

import (
	"goblog/common"
	"goblog/router"
	"log"
	"net/http"
)

func init() {
	// 模板加载
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println("启动服务出错：", err)
	}
}
