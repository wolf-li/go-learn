package main
import (
	"fmt"
	"errors"
)

func Parent() error {
	err := method()
	return err
}

func method() error {
	return errors.New("有问题")
}

func div(a, b int)(res int ,err error){
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	res = a/b
	return
}

func server()(res int, err error){
	res, err = div(1, 3)
	if err != nil {
		return
	}

	// 正常逻辑
	res++
	
	return 
}

func main(){
	res, err := server()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}