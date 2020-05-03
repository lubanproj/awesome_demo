package main

import (
	"fmt"
	"net"
)

func call() (string, error) {
	// 第一步，建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		return "", err
	}

	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	// 第二步，发送请求包
	if _, err := conn.Write([]byte("hello")); err != nil {
		fmt.Println("conn Write error,", err)
		return "", err
	}

	// 第三步，接收 server 回包
	buffer := make([]byte, 1024)
	recvNum, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("conn Read error,", err)
		return "", err
	}

	msg := string(buffer[:recvNum])
	fmt.Println("recv msg from server : ", msg)

	return msg, nil
}