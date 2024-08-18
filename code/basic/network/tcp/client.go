package main

import (
	"net"
	"fmt"
)

func main() {
	conn, _ := net.DialTCP("tcp", nil,
		&net.TCPAddr{
			IP : net.ParseIP("127.0.0.1"),
			Port: 8000,
		})	
	defer conn.Close()
	var a string
	for {
		fmt.Scanln(&a)
		if a == "q" {
			break
		}
		conn.Write([]byte(a))
	}

}