package main

import (
	"io"
	"fmt"
	"os"
	"log"
)

func main(){
	// 一次性写
	// err := os.WriteFile("hello", []byte("Hello, Gophers!\n"), 0000)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 使用 os.OpenFile 函数
	file, err := os.OpenFile(os.Args[1], os.O_CREATE| os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	file.Write([]byte("你好"))

	file1, err := os.OpenFile(os.Args[1], os.O_CREATE| os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file1.Close()
	byteData, err := io.ReadAll(file1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(byteData))
}