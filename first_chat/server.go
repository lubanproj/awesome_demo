package first_chat

import (
	"fmt"
	"net"
)

func main() {
	lis, err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.Accept()
		defer conn.Close()

		if err != nil {
			fmt.Println("listener accept err, ", err)
			continue
		}

		buffer := make([]byte, 1024)
		recvNum, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("conn Read err, ", err)
			continue
		}

		msg := string(buffer[:recvNum])
		fmt.Println("recv msg : ", msg)
	}
}