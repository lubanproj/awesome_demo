package first_chat

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	if _, err := conn.Write([]byte("hello")); err != nil {
		fmt.Println("conn Write error,", err)
		return
	}

	buffer := make([]byte, 1024)
	recvNum, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("conn Read error,", err)
	}

	msg := string(buffer[:recvNum])
	fmt.Println("recv msg from server : ", msg)
}