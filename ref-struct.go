package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct{
	Age int
	Name string `json:name`
}

func refSetUser(obj any){
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()

	for i := 0; i < t.NumField(); i++ {
		filed := v.Field(i)
		jsonFiled := t.Field(i).Tag.Get("json")
		// if filed.Type.Kind() != reflect.String{
		// 	continue
		// }
		if jsonFiled == "" {
			continue
		}
		filed.SetString(strings.ToUpper(filed.String()))
	}
}

func main() {
	u1 := User{
		Name: "x",
		Age: 12,
	}
	t := reflect.TypeOf(u1)
	v := reflect.ValueOf(u1)

	// NumField() 结构体属性数量
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		jsonFiled := filed.Tag.Get("json")
		if jsonFiled == "" {
			jsonFiled = filed.Name
		}
		fmt.Printf("name: %s type: %s json: %s value: %v\n", filed.Name, filed.Type, jsonFiled, v.Field(i))
	}

	refSetUser(&u1)
	fmt.Println(u1)
}