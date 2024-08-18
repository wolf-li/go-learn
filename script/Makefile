APP := $(shell basename ${PWD})
AMD64=GOARCH=amd64
ARM64=GOARCH=arm64

build:  
	go build -o ${APP} main.go  
  
run:  
	go run -race main.go  
  
clean:  
	go clean  

release-amd64:
	CGO_ENABLED=0 GOOS=linux ${AMD64} go build -o ${APP}
	tar czf ${APP}-linux-amd64.tar.gz ./${APP}
	go clean
	CGO_ENABLED=0 GOOS=windows ${AMD64} go build -o ${APP}
	tar czf ${APP}-windows-amd64.tar.gz ./${APP}.exe
	go clean
	CGO_ENABLED=0 GOOS=darwin ${AMD64} go build -o ${APP}
	tar czf ${APP}-darwin-amd64.tar.gz ./${APP}
	go clean

  

# https://colynn.github.io/2020-03-03-using_makefile/