# grpc

## 单体架构
1. 只能对整体进行扩容
2. 相互调用
3. 耦合程度高

## 微服务架构
可以按照服务进行扩容
各服务之间相互独立，可以独立部署

问题：
1. 代码冗余
2. 服务和服务之间存在调用关系

RPC（远程过程调用）
grpc（框架）
gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

优势：
性能高相对于 json 

RPC 实现需要解决的问题
函数调用
服务端定义 IDL 文件，编译工具生成 sub 桩文件，（protocol buffer pb文件实现）相当于生成了静态库，实现函数映射
网络层面屏蔽

## 环境准备
工具 protoc
```
// linux
$ apt install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+
// macos
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```
---
https://grpc.io/docs/protoc-installation/

Go plugins for the protocol compiler:
Install the protocol compiler plugins for Go using the following commands:
```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
Update your PATH so that the protoc compiler can find the plugins:
`$ export PATH="$PATH:$(go env GOPATH)/bin"`

---
https://grpc.io/docs/languages/go/quickstart/
