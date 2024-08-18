package main

import "fmt"

func timer() func() int {
	count := 0
	return func() int{
		count++
		return count
	}
}

func main(){
	counter := timer()
	fmt.Println(counter())
	fmt.Println(counter())
}