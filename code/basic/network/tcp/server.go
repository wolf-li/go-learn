package main

import (
	"net"
	"fmt"
	"io"
)


func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp","127.0.0.1:8000")
	listern, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, err := listern.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 获取客户端地址
		fmt.Println(conn.RemoteAddr().String()," connect")
		for {
			var buf []byte = make([]byte, 1024)
			n, err := conn.Read(buf)
			if err == io.EOF{
				fmt.Println(conn.RemoteAddr().String()," disconnect")
				break
			}
			fmt.Println(string(buf[0:n]))
		}
	}
}