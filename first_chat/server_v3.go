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


		go handleConnV2(conn)
	}

}

func handleConnV2(conn net.Conn) {

	for {
		buffer := make([]byte, 1024)
		recvNum , err := conn.Read(buffer)
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
