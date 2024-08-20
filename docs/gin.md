# gin 框架
官方文档：https://gin-gonic.com/docs/

快速使用
```go

```

初始化配置
关闭 debug 日志
```go
gin.SetMode("release")  // 没有debug 日志
```
启动方式
第一种：
```go
router.Run(":8000")
```
第二种：
```go
http:ListenAndServer(":8000",router)
```

## 响应
http 状态码

返回字符串
```go
    r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.String(http.Status.OK, "返回 hello")
	})
```

返回 json
```go
    r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
```

返回 xml
```go
func returnXML(c *gin.Context){
	c.XML(http.StatusOK, gin.H{"user":"lala","msg":"wx","status":http.StatusOK})
}
```

返回 yml
```go
func returnYAML(c *gin.Context){
	c.YAML(http.StatusOK, gin.H{"user":"lala","msg":"wx","status":http.StatusOK})
}
```

返回 html
先使用 LoadHTMLGlob() 或 LoadHTMLFiles() 方法加载模版
```go
func returnHTML(c *gin.Context){
	// gin.H{} 向 html 模版传递参数
	c.HTML(200, "index.html", gin.H{"body":"return html"})
}

// r.LoadHTMLGlob("./*.html")
r.LoadHTMLFiles("index.html")
r.GET("/html", returnHTML)
```

文件响应
go 制定文件目录是相对项目的
使用 router.StaticFS() 函数 或 StaticFile() 函数
```go
r.StaticFile("/static", "index.html")
```

重定向
```go
func rewrite(c *gin.Context){
	// 永久转移，可以被浏览器缓存
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}
r.GET("/rewrite", rewrite)
```

## 请求参数
获取 querystring 参数，请求 url /user/search?username=小王子&address=沙河 
```go
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
router.GET("/user/search", querystring)
```
获取 form 表单请求 
```go
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
router.POST("/user/search", queryform)
```

获取 JSON 参数 请求 POST /user/searcch
```go
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
router.POST("/user/search", queryjson)
router.PUT("/user/search", queryjson)
```

获取 path 参数 请求的参数通过 url 传递例子 /user/search/小王子/沙河  (不需要校验)
```go
func querypath(ctx *gin.Context){
	username := ctx.Param("username")
	address := ctx.Param("address")
	ctx.JSON(http.StatusOK, gin.H{
		"msg":"ok", 
		"username":username,
		"address":address,
	})
}
```

### 参数绑定
为了方便获取请求相关参数，提高开发效率，可以基于请求的 content-type 识别请求数据类型并利用反射机制自动提取请求中的 QueryString 、form 表单、JSON、 XML 等参数结构体。利用 shouldbind()
公共代码
```go
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
```
// 利用 shouldbind 进行获取 POSt JSON
```go
func bindJson(ctx *gin.Context){
	var ujson User
	if err := ctx.ShouldBind(&ujson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": ValidUser(&ujson).Error()})
	}else{
		ctx.JSON(http.StatusOK, ujson)
	}
}
router.POST("/user/search", bindJson)
```

// 利用 shouldbind 获取 form 表单参数
```go
func bindForm(ctx *gin.Context){
	var uform User
	if err := ctx.ShouldBind(&uform); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}else{
		ctx.JSON(http.StatusOK, uform)
	}
}
router.POST("/user/search", bindForm)
```

// 利用 shouldbind 获取 QueryString GET /user/search?username=q1mi&address=123456
```go
func bindQuery(ctx *gin.Context){
	var uQuery User
	if err := ctx.ShouldBind(&uQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}else{
		ctx.JSON(http.StatusOK, uQuery)
	}
}
router.GET("/user/search", bindQuery)
```
## 文件上传
前端页面
```go
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>上传文件示例</title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="f1">
    <input type="submit" value="上传">
</form>
</body>
</html>
```
restful 风格 api 
URL POST /api/files
不要使用 Thunder 进行测试 （连接远程主机情况下）
```go
func files(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(file.Filename)
	dest := "./" + file.Filename
	ctx.SaveUploadedFile(file, dest)
	ctx.JSON(http.StatusOK, gin.H{"message": "save file: "+file.Filename})
}
```
服务端保存文件的方式
SaveUploadedFile
```go
c.SaveUploadedFile(file, dst)  // 文件对象  文件路径，注意要从项目根路径开始写
```
Create+Copy
file.Open的第一个返回值就是我们讲文件对象中的那个文件（只读的），我们可以使用这个去直接读取文件内容
```go
file, _ := c.FormFile("file")
log.Println(file.Filename)
// 读取文件中的数据，返回文件对象
fileRead, _ := file.Open()
dst := "./" + file.Filename
// 创建一个文件
out, err := os.Create(dst)
if err != nil {
  fmt.Println(err)
}
defer out.Close()
// 拷贝文件对象到out中
io.Copy(out, fileRead)
```

读取上传的文件
```go
file, _ := c.FormFile("file")
// 读取文件中的数据，返回文件对象
fileRead, _ := file.Open()
data, _ := io.ReadAll(fileRead)
fmt.Println(string(data))
```
这里的玩法就很多了
例如我们可以基于文件中的内容，判断是否需要保存到服务器中

上传多个文件
```go
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

router.POST("/files", multifiles)
```

## 重定向
http 重定向
redirect
```go
func redirect(ctx *gin.Context){
	// 临时转移 307
	// ctx.Redirect(http.StatusTemporaryRedirect, "https://www.baidu.com")
	// 永久转移 308
	ctx.Redirect(http.StatusPermanentRedirect, "https://www.baidu.com")
}
router.GET("/1", redirect)
```
在Gin框架中，直接修改ctx.Request.URL.Path并尝试通过router.HandleContext(ctx)重新路由是不推荐的，因为这可能会导致一些不预期的行为，特别是与中间件和请求上下文相关的状态可能不会被正确重置或更新。
不过，如果你确实需要在某个路由处理函数中“重定向”到另一个路由处理函数，你可以通过定义独立的函数来实现这一点，然后在需要的路由中调用这些函数。但请注意，这并不是真正的路由重定向，而是函数调用。
路由重定向
```go
router.PUT("/2", func(ctx *gin.Context) {
	ctx.Request.URL.Path = "/1"
	router.HandleContext(ctx)
})
```

## 路由
普通路由
```go
router.GET("/index", func(ctx *gin.Context){...})
// 请求所有方法
router.Any("/index", func(ctx *gin.Context){...})
// 统一处理未知路由
router.NoRoute(func(ctx *gin.Context){
	ctx.HTML(http.StatusNotFound, "views/404.html", nil)
})
```
路由组
可以将有共同 url 前缀的路由划分的到一个路由组，习惯上用{} 包裹共同路由
```go
router := gin.Default()
userGroup := router.Group("/user")
{
	userGroup.POST("/login", func(ctx *gin.Context){...})
	userGroup.GET("/logout", func(ctx *gin.Context){...})
	userGroup.POST("/register", func(ctx *gin.Context){...})
}
bookGroup := router.Group("/book")
{
	bookGroup.GET("/search",  func(ctx *gin.Context){...})
	bookGroup.PUT("/update",  func(ctx *gin.Context){...})
	bookGroup.POST("/add",  func(ctx *gin.Context){...})
	bookGroup.DELETE("/delete",  func(ctx *gin.Context){...})
}
```

## 中间件
Gin 框架允许开发者在处理请求的过程中，加入用户自己的钩子 hook函数，这个钩子函数就叫做中间件，中间件适合处理一些公共的业务逻辑，比如登陆认证、权限校验、数据分页、记录日志、耗时统计等
定义中间件,中间件定义必须是 gin.HandlerFunc 类型
记录接口耗时的中间件
```go
func StatCost() gin.HandlerFunc{
	return func(c *gin.Context){
		start := time.Now()
		c.Set("name":"xiao")
		c.Next()
		cost := time.Since(start)
		log.Println(cost)
	}
}
```


---
参考文档：
https://www.liwenzhou.com/posts/Go/gin/#c-0-6-2
