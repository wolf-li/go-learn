# GUI 开发

fnye GUI 开发框架
特点：跨平台
运行需要：
1. Go 版本低 1.13 建议单独环境开发
2. C 编译器 
3. 系统图形驱动程序

linux 下
```bash
sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev
```

项目创建步骤
```go
cd myapp
go mod init MODULE_NAME
go get fyne.io/fyne/v2
go mod tidy
```
[官方文档](https://go-circle.cn/fyne-press/v1.0/1-getting-started/introduction.html)
