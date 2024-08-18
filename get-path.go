package main

import (
	"fmt"
	"runtime"
)

func GetCurrentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	return file
  }

func main(){
	fmt.Println(GetCurrentFilePath)
}