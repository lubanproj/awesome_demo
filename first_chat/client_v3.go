package main

import (
	"fmt"
	"net"
)

var globalConn net.Conn
var err error

func init() {
	globalConn, err = net.Dial("tcp", "127.0.0.1:8000")
}

func getGlobalConn() (net.Conn, error) {

	if globalConn != nil {
		return globalConn, nil
	}

	conn, err := net.Dial("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("net Dial error, ", err)
		return nil , err
	}

	return conn, nil
}


func callV3() (string, error) {

	var err error

	globalConn, err = getGlobalConn()
	if err != nil {
		fmt.Println("getGlobalConn error, ", err)
		globalConn, err = getGlobalConn()
	}

	defer func() {
		if err != nil && globalConn != nil{
			globalConn.Close()
			globalConn, err = getGlobalConn()
		}
	}()

	// 第二步，发送请求包
	if _, err := globalConn.Write([]byte("hello")); err != nil {
		fmt.Println("conn Write error,", err)
		return "", err
	}

	// 第三步，接收 server 回包
	buffer := make([]byte, 1024)
	recvNum, err := globalConn.Read(buffer)
	if err != nil {
		fmt.Println("conn Read error,", err)
		return "", err
	}

	msg := string(buffer[:recvNum])
	fmt.Println("recv msg from server : ", msg)

	return msg, nil
}
