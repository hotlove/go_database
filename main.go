package main

import (
	"fmt"
	_ "go_database/entity"
	"go_database/server/socket"
	"log"
	"net"
	"net/http"
	"os"
	"github.com/golang/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
		msg := "Received: " + reply
		fmt.Println("Sending to client: " + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
func main() {
	// 启动默认http服务
	//http.Start()
	//http.NewServer().Start("127.0.0.1", 8080)

	//testIp()

	//startSocketServer()
	//socket.NewSocketClient()

	// websocket
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func startSocketServer() {
	socket.StartServer()
}

func testIp() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
