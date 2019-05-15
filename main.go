package main

import (
	_ "go_database/entity"
	"go_database/server/http"
)

func main() {
	// 启动默认http服务
	http.Start()
	//http.NewServer().Start("127.0.0.1", 8080)
}
