package main

import (
	"fmt"
	_ "go_database/entity"
	"go_database/server/socket"
	"net"
	"os"
)

func main() {
	// 启动默认http服务
	//http.Start()
	//http.NewServer().Start("127.0.0.1", 8080)

	//testIp()

	//startSocketServer()
	socket.NewSocketClient()
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
