#!/bin/bash

if [[ $# -ne 1 ]];then
    echo "usage: ./initFyne program"
    exit 1
fi
mkdir $1
cd $1
cat > main.go < EOF
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
EOF
go mod init $1
go get fyne.io/fyne/v2
go mod tidy