package main

import "fmt"

func add(a int){
	fmt.Printf("func inside: %p\n",&a)
	a += a
	fmt.Println(a)
}

func add1(a *int){
	fmt.Printf("func inside: %p\n",a)
	*a = *a + *a
	fmt.Println(*a)
}

func main(){
	var a int = 3
	fmt.Printf("func outside: %p\n",&a)
	add1(&a)
	fmt.Println(a)
}