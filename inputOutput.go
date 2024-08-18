package main

import "fmt"

func main(){
	// fmt.Println("hello","world")
	// fmt.Println("你好")

	// fmt.Printf("%s\n","你好")
	// fmt.Printf("%v %T\n", "你好","你好")
	// fmt.Printf("%d\n", 123)
	// fmt.Printf("%.2f\n", 234.1234)
	// fmt.Printf("%#v\n","")

	// input
	fmt.Println("input your name")
	var name string
	fmt.Scan(&name)
	fmt.Println("your name is: ",name)
}