package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WebSocketServer struct {
	Upgrader websocket.Upgrader
}

func initWebSocket() *WebSocketServer {
	// 配置websocket服务
	upgrader := websocket.Upgrader{
		// 读取存储空间大小
		ReadBufferSize: 1024,
		// 写入存储空间大小
		WriteBufferSize: 1024,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &WebSocketServer{
		Upgrader: upgrader,
	}
}
