package socket

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type SocketServer struct {
}

var connManager = map[string]net.Conn{}

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
		fmt.Println("wait connect...")
		conn, err := listener.Accept()

		if err != nil {
			checkError(err)
			continue
		}

		// 开启一个goruntine 处理链接
		go registerConn(conn)

	}
}

func registerConn(conn net.Conn) {

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		checkError(err)
		identity := string(buffer[:n])

		// 解析参数
		//  sender,receiver,message
		messages := strings.Split(identity, ",")
		var sender = messages[0]
		var receiver = messages[1]
		var value = messages[2]
		// 说明是注册链接信息
		if value == "r" {
			if sender != "" {
				// 如果没有注册链接则注册一个链接
				if _, ok := connManager[sender]; !ok {
					connManager[sender] = conn
					fmt.Println(sender+"register conn.....", conn)
					conn.Write([]byte("register successfully"))
				}
			}
		} else {
			// 说明是发送消息

			//senderConn, ok := connManager[sender]
			//if !ok {
			//	// 说明没有注册信息抛弃
			//	log.Println(sender + "not register conn")
			//}

			receiverConn, ok := connManager[receiver]
			if !ok {
				// 说明接受者没注册
				log.Println(receiver + "not register conn")
			}

			//fmt.Printf("sender %v \n", senderConn)
			//fmt.Printf("receiver %v \n", receiverConn)

			// 从发送者收到数据
			fmt.Println(sender + "to" + receiver + " send message: [" + value + "]")

			// 转发给接收者
			receiverConn.Write([]byte(value))

		}
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
