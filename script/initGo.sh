#!/bin/bash
# version: 0.0.1
# auth: li
# date: 2024-08-14
# description: init golang program 
# usage: ./initGo.sh programname


if [[ $# -ne 1 ]];then
    echo "usage: ./initGo.sh programname"
    echo "wrong args"
    exit 1
fi

mkdir $1
cd $1
go mod init $1
cat > main.go << EOF
package main

import (
    "fmt"
)

func ErrHandler(err error){
	if err != nil {
		log.Fatal(err)
		return 
	}
}

func main(){
    fmt.Println("hello world")
}
EOF

go build
chmod o+x $1
./$1
rm -f $1