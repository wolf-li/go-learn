package main
import "fmt"

type Code int

const (
	SuccessCode Code = 0
	InvalidCode Code = 10003
	FailCode Code = 11111
)

func (c Code)GetMsg()(string){
	switch c{
	case SuccessCode:
		return "成功"
	case InvalidCode:
		return "无效参数"
	case FailCode:
		return "失败操作"
	default:
		return "未知"
	}
}

func (c Code) CodeStatus()(Code, string){
	return c,c.GetMsg()
}

func main(){
	var c1 Code = 1
	fmt.Println(c1.CodeStatus())
}