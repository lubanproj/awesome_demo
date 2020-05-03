package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	for {
		conn , err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer conn.Close()
		if err != nil {
			panic(err)
		}


		go handleConnV4(conn)
	}

}

func handleConnV4(conn net.Conn) {

	for {
		// 修改 conn.Read 为 io.ReadFull
		buffer := make([]byte, 5)
		recvNum , err := io.ReadFull(conn, buffer)
		if err == io.EOF {
			// client 连接关闭
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}

		msg := string(buffer[:recvNum])
		fmt.Println("recv from client: ",msg)

		handler()

		conn.Write([]byte("world"))
	}

}
