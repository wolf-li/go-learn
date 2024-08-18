package main

import "fmt"

type People struct{
	Name string
}

func (p People) printInfo(){
	fmt.Printf("name: %s\n", p.Name)
}

type Student struct{
	People
	Age int
}

func (s Student) printInfo(){
	fmt.Printf("name: %s, age: %d\n", s.Name, s.Age)
}

func main(){
	p := People{
		Name:"lala",
	}
	p.printInfo()
	s1 := Student{
		People: People{Name:"qlql"},
		Age:12,
	}
	s1.printInfo()
}