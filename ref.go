package main

import (
	"fmt"
	"reflect"
)

func getType(obj any){
	t := reflect.TypeOf(obj)
	fmt.Println(t, t.Kind())
	switch t.Kind() {
	case reflect.Float32:
		fmt.Println("float32")
	case reflect.Int:
		fmt.Println("int")
	case reflect.Slice:
		fmt.Println("slice")
	case reflect.String:
		fmt.Println("string")
	case reflect.Struct:
		fmt.Println("struct")
	}
}

func refValue(obj any){
	value := reflect.ValueOf(obj)
	fmt.Println(value, value.Type())
	switch value.Kind(){
	case reflect.Int:
		fmt.Println(value.Int())
	case reflect.String:
		fmt.Println(value.String())
	}
}

func refSetValue(obj any){
	objValueOf := reflect.ValueOf(obj)
	elem := objValueOf.Elem()
	switch elem.Kind() {
	case reflect.String:
		elem.SetString("haha")
	}
}

func main() {
	name := "lala"
	getType(name)
	getType(struct{Name string}{Name:"xix"})
	refValue("lala")
	refValue(3)

	refSetValue(&name)
	fmt.Println(name)
}