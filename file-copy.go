package main

import (
	"fmt"
	"io"
	"os"
	"log"
)

func errHandler(err error){
	if err != nil {
		log.Fatal(err)
		return 
	}
}

func main(){
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s sourcefile destfile\n", os.Args[0][2:])
		return
	}
	
	read, err := os.Open(os.Args[1])
	errHandler(err)
	defer read.Close()

	write, err := os.Create(os.Args[2])
	errHandler(err)
	defer write.Close()

	n, err := io.Copy(write, read) 
	errHandler(err)
	
	fmt.Println(n)
}