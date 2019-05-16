package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type SocketServer struct {
}

func StartServer() {
	// ip + port 地址
	service := "127.0.0.1:12200"

	// 获取一个tcp地址对象
	tcpAdder, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 开始监听ip+port 并返回一个链接
	listener, err := net.ListenTCP("tcp", tcpAdder)
	checkError(err)
	for {
		// 阻塞等待链接过来
		conn, err := listener.Accept()
		fmt.Println("connect coming...")
		if err != nil {
			checkError(err)
			continue
		}

		// 开启一个goruntine 处理链接
		go handleTCPClient(conn)

	}
}

func handleTCPClient(conn net.Conn) {
	defer conn.Close()

	// 循环读取数据
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		checkError(err)
		fmt.Println("client message:" + string(buffer[:n]))

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
		}

		message := scanner.Text()
		fmt.Println("send message..." + message)
		conn.Write([]byte(message))

	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}
