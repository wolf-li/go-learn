package main

import (
	"fmt"
	"reflect"
)

type Student struct{
	Name string
	Age int
}

func (Student) See(name string){
	fmt.Println("see name:",name)
}

func main() {
	s := Student{
		Name: "小智",
		Age: 21,
	}

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumMethod(); i++ {
		methodType := t.Method(i)
		fmt.Printf("name: %s, type: %s\n",methodType.Name, methodType.Type)
		methodValue := v.Method(i)
		methodValue.Call([]reflect.Value{
			reflect.ValueOf("lala"),
		})
	}
}