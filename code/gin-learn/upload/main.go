package main

import (
	// "fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

// 单文件上传
func files(ctx *gin.Context) {
	uploadfile, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	dest := "./" + uploadfile.Filename
	ctx.SaveUploadedFile(uploadfile, dest)
	ctx.JSON(http.StatusOK, gin.H{
		"filename": uploadfile.Filename,
		"filesize": uploadfile.Size,
		"createTime": time.Now(),
	})
}

// 多文件上传
func multifiles(ctx *gin.Context){
	multifile, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg":err.Error()})
		return
	}
	
	files := multifile.File["mult"]
	for _, file := range files{
		dest := "./" + file.Filename
		ctx.SaveUploadedFile(file, dest)
	}
	ctx.JSON(http.StatusOK, gin.H{"msg":"ok"})
}

// 重定向  
func redirect(ctx *gin.Context){
	// 临时转移 307
	// ctx.Redirect(http.StatusTemporaryRedirect, "https://www.baidu.com")
	// 永久转移 308
	ctx.Redirect(http.StatusPermanentRedirect, "https://www.baidu.com")
}

// 路由重定向


func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./index.html")
	router.MaxMultipartMemory = 8 << 20
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	
	router.POST("/files", multifiles)
	router.GET("/1", redirect)
	router.POST("/1", redirect)
	router.PUT("/1", redirect)
	router.PUT("/2", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/1"
		router.HandleContext(ctx)
	})
	router.Run(":9090")
}
