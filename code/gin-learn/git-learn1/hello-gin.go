package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context){
	c.JSON(200,gin.H{
		"msg":"ok",
	})
}

func returnString(c *gin.Context){
	c.String(http.StatusOK, "return string")
}

func returnXML(c *gin.Context){
	c.XML(http.StatusOK, gin.H{"user":"lala","msg":"wx","status":http.StatusOK})
}

func returnYAML(c *gin.Context){
	c.YAML(http.StatusOK, gin.H{"user":"lala","msg":"wx","status":http.StatusOK})
}

func returnHTML(c *gin.Context){
	c.HTML(200, "index.html", gin.H{"body":"return html"})
}

func rewrite(c *gin.Context){
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}

func getid(c *gin.Context){
	fmt.Println(c.Param("id"))
	c.JSON(200, gin.H{"status":"ok"})
}

func main() {
	// 初始化
	gin.SetMode("release")  // 没有debug 日志
	r := gin.Default()
	// 注册路由
	r.GET("/json", index)
	r.GET("/string", returnString)
	r.GET("/xml", returnXML)
	// 下载 yaml 文件
	r.GET("/yaml", returnYAML)

	// r.LoadHTMLGlob("./*.html")
	r.LoadHTMLFiles("index.html")
	r.GET("/html", returnHTML)

	r.StaticFile("/static", "index.html")
	r.GET("/rewrite", rewrite)

	// 获取请求值
	r.GET("/value/:id", getid)
	r.PUT("/pvalue/:id", getid)
	// 绑定端口
	r.Run(":8090")
}