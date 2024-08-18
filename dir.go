package main

import (
	"os"
	"fmt"
	"log"
)
func ErrHandler(err error){
	if err != nil {
		log.Fatal(err)
		return 
	}
}
func main(){
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s filename\n", os.Args[0][2:])
		return
	}
	dir, err := os.ReadDir(os.Args[1])
	ErrHandler(err)

	for _, entry := range dir{
		info, _ := entry.Info()
		fmt.Println(entry.IsDir(), entry.Name(), info.Size())
	}
}