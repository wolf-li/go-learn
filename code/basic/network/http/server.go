package main

import (
	"fmt"
	"net/http"
)

func Index(write http.ResponseWriter, request *http.Request){
	fmt.Println(request.URL.Path)
	write.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/index", Index)
	fmt.Println("web address: http://127.0.0.1:8000")
	http.ListenAndServe("127.0.0.1:8000",nil)
}