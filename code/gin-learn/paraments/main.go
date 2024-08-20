package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" form:"username" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
}

func ValidUser(u *User) error {
	// 使用反射获取值判断
	if u.Address == ""{
		return fmt.Errorf("address 参数为空")
	}
	if u.Username == ""{
		return fmt.Errorf("username 参数为空")
	}
	return nil
}

// 利用 shouldbind 进行获取 POSt JSON
func bindJson(ctx *gin.Context){
	var ujson User
	if err := ctx.ShouldBind(&ujson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": ValidUser(&ujson).Error()})
	}else{
		ctx.JSON(http.StatusOK, ujson)
	}
}

// 利用 shouldbind 获取 form 表单参数
func bindForm(ctx *gin.Context){
	var uform User
	if err := ctx.ShouldBind(&uform); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}else{
		ctx.JSON(http.StatusOK, uform)
	}
}

// 利用 shouldbind 获取 QueryString GET /user/search?username=q1mi&address=123456
func bindQuery(ctx *gin.Context){
	var uQuery User
	if err := ctx.ShouldBind(&uQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}else{
		ctx.JSON(http.StatusOK, uQuery)
	}
}

// 获取 querystring 参数，请求 GET url /user/search?username=小王子&address=沙河
func querystring(ctx *gin.Context) {
	username := ctx.Query("username")
	address := ctx.Query("address")
	if username == "" || address == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "missing args"})
	}else{
		ctx.JSON(http.StatusOK, gin.H{
			"msg":"ok", 
			"username":username,
			"address":address,
		})
	}
}

// 获取 form 参数 请求 POST url /user/search
func queryform(ctx *gin.Context){
	username := ctx.PostForm("username")
	address := ctx.PostForm("address")
	if username == "" || address == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "missing args"})
	}else{
		ctx.JSON(http.StatusOK, gin.H{
			"msg":"ok", 
			"username":username,
			"address":address,
		})
	}
}
// 获取 JSON 参数 请求 POST /user/searcch
func queryjson(ctx *gin.Context){
	// 从 ctx.Request.Body 中获取数据,
	b, _ := ctx.GetRawData()
	var m map[string]any
	// 反序列化数据
	if err := json.Unmarshal(b, &m); err != nil || len(m) == 0 {
		fmt.Println("没有数据")
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"没有参数"})
		return
	}
	ctx.JSON(http.StatusOK, m)
}

// 获取 path 参数 请求的参数通过 url 传递例子 /user/search/小王子/沙河 
func querypath(ctx *gin.Context){
	username := ctx.Param("username")
	address := ctx.Param("address")
	ctx.JSON(http.StatusOK, gin.H{
		"msg":"ok", 
		"username":username,
		"address":address,
	})
}


func main() {
	router := gin.Default()
	// router.GET("/user/search", querystring)
	// router.POST("/user/search", queryform)
	// router.PUT("/user/search", queryjson)
	// router.GET("/user/search/:username/:address", querypath)

	// router.POST("/user/search", bindJson)
	router.POST("/user/search", bindForm)
	router.GET("/user/search", bindQuery)

	router.Run(":9090")
}
