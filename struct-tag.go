package main

import (
	"fmt"
	"encoding/json"
)

type Student struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Male string `json:"male"`
	Password string `json:"-"`
}

func main(){
	user := Student{Name:"rase", Age:12, Male:"m",Password:"123"}
	byteDate, _ := json.Marshal(user)
	fmt.Println(string(byteDate))
}