package main

import (
	"fmt"
	_ "go_database/entity"
	"go_database/server/socket"
	"net"
	"os"
	"time"
)

func main() {
	// 启动默认http服务
	//http.Start()
	//http.NewServer().Start("127.0.0.1", 8080)

	//testIp()

	//startSocketServer()
	//socket.NewSocketClient()

	// websocket
	//websocket.StartServer()

	// redis
	//value := util.RedisGetValue("foo")
	//fmt.Println(value)

	// 定时器
	tick := time.NewTicker(1 * time.Second)

	for {
		select {
		case dateTime := <-tick.C:
			// 这里面有格式化时间在里面
			fmt.Println(dateTime.Format("2006/01/02 15/04/05"), "执行了")
		}
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
