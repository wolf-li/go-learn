package main

import (
	"fmt"
	"os"
)

func main(){
	byteData, _ := os.ReadFile("main.go")
	fmt.Println(string(byteData))
}