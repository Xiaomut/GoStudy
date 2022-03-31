package main

import (
	"fmt"
	"net"
)

// server 端

func main() {
	// 1. 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:6666")	
	if err != nil{
		fmt.Println("start server failed", err)
	}

	// 2. 等待客户建立链接
	
}