package main

import "fmt"

type Animal interface{
	sing()
}

type Chicken struct{
	Name string
}

func (ck Chicken)sing(){
	fmt.Printf("%s 在唱歌\n",ck.Name)
}

type Cat struct{
	Name string
}

func (c Cat)sing(){
	fmt.Println("唱歌")
}

func (c Cat)jump(){
	fmt.Println("跳")
}

func (c Cat)rap(){
	fmt.Println("rap")
}

func sing(obj Animal){
	switch obj.(type){
	case Chicken:
		fmt.Println("鸡")
	case Cat:
		fmt.Println("猫")
	}
	obj.sing()
}

func main(){
	chicken1 := Chicken{Name:"caixukun"}
	sing(chicken1)
	cat1 := Cat{Name:"mimi"}
	sing(cat1)
}