package first_chat

import (
	"fmt"
	"net"
)

func main() {
	// 第一步，监听请求
	lis, err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	for {
		// 第二步，获取一个连接
		conn, err := lis.Accept()
		defer conn.Close()

		if err != nil {
			fmt.Println("listener accept err, ", err)
			continue
		}

		// 第三步，读取 client 请求包
		buffer := make([]byte, 1024)
		recvNum, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("conn Read err, ", err)
			continue
		}

		// 第四步，发送响应包
		msg := string(buffer[:recvNum])
		fmt.Println("recv msg : ", msg)
	}
}