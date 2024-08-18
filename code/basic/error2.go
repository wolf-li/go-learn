package main

import (
	// "errors"
	"fmt"
	"os"
)

func init(){
	_, err := os.ReadFile("xx")
	if err != nil {
		panic(err.Error())
	}
}

func main(){
	fmt.Println("main")
}