package websocket

import (
	"log"
	"net/http"
)

func StartServer() {
	// 当有请求访问ws时，执行此回调方法
	http.HandleFunc("/ws", wsHandler)

	// 监听127.0.0.1:7777
	log.Println("websocket 服务启动 port:7777")
	err := http.ListenAndServe("0.0.0.0:7777", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}
}
