package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://127.0.0.1:8000/index")
	if err != nil {
		fmt.Println(err)
		return
	}
	byteData, _ := io.ReadAll(response.Body)
	fmt.Println(string(byteData))
}