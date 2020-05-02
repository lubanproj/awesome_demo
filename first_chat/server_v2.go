package main

import (
	"fmt"
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


		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	recvNum , err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}

	msg := string(buffer[:recvNum])
	fmt.Println("recv from client: ",msg)

	conn.Write([]byte("world"))
}