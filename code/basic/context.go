package main

import (
	"context"
	"fmt"
)

func GetUser(ctx context.Context){
	fmt.Println(ctx.Value("name"))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "lala")
	GetUser(ctx)
}