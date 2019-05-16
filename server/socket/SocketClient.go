package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func NewSocketClient() {
	service := "127.0.0.1:12200"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)

	handleChat(conn)
}

func handleChat(conn net.Conn) {

	registerInfo := "2,1,r"
	// 第一登陆注册链接
	conn.Write([]byte(registerInfo))

	go func() {
		for {
			buffer := make([]byte, 2048)
			n, err := conn.Read(buffer)
			checkError(err)
			fmt.Println("receiver message:" + string(buffer[:n]))
		}
	}()

	for {
		// 等待键盘输入数据
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

func handleClient(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
		}

		message := scanner.Text()
		fmt.Println("send message..." + message)
		conn.Write([]byte(message))

		n, err := conn.Read(buffer)
		checkError(err)
		fmt.Println("server message..." + string(buffer[:n]))

	}
}
