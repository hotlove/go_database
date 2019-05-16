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

	tcpConn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)

	buffer := make([]byte, 2048)
	for {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
		}

		message := scanner.Text()
		fmt.Println("send message..." + message)
		tcpConn.Write([]byte(message))

		n, err := tcpConn.Read(buffer)
		checkError(err)
		fmt.Println("server message..." + string(buffer[:n]))

	}

	os.Exit(0)

}
