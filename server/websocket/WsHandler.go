package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 获取websocketserver 对象
var websocketServer = initWebSocket()

// 用于管理连接
var connManager = map[int]*websocket.Conn{}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	// 完成http应答，在httpheader中放下如下参数
	conn, err := websocketServer.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		panic(err)
		return // 获取连接失败直接返回
	}

	// for循环负责接收链接
	for {
		// 开启协程处理连接 一个协程处理一个连接
		go connectionHandler(conn)

	}
}

func connectionHandler(conn *websocket.Conn) {
	// 捕获异常
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			closeConn(conn)
		}
	}()

	// 循环从连接中读取数据
	for {
		// 只能发送Text, Binary 类型的数据,下划线意思是忽略这个变量.
		_, data, err := conn.ReadMessage()

		packet := &Packet{}

		// 解析json数据到结构体对象
		json.Unmarshal(data, packet)

		// 建立连接请求
		if packet.Type == MESSAGE_TYPE_LOING {

			registerConnHandler(packet.SendProfileId, conn)
		}

		// 发送消息
		if packet.Type == MESSAGE_TYPE_COMMON {
			chatHandler(packet)
		}

		if packet.Type == MESSAGE_TYPE_LOGOUT {
			disconnectHandler(packet.SendProfileId)
		}
		if err != nil {
			panic(err)
		}
	}
}

// 注册连接
func registerConnHandler(profileId int, conn *websocket.Conn) {
	connManager[profileId] = conn
}

// 处理聊天消息
func chatHandler(packet *Packet) {
	// 找到接收人连接
	recvConn := connManager[packet.RecvProfileId]

	// 发送消息给到接收者
	writerErr := recvConn.WriteMessage(websocket.TextMessage, []byte(packet.Message))
	if writerErr != nil {
		log.Println(writerErr)
	}
}

func disconnectHandler(profileId int) {
	// 从管理map中移除该连接
	delete(connManager, profileId)
}
func closeConn(conn *websocket.Conn) {
	conn.Close()
}
