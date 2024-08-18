package main

import "testing"

func Test_hello(t *testing.T) {
	if ans := hello(); ans != "hello"{
		t.Errorf("return not hello")
	}

	// 子测试
	// tests := []struct {
	// 	name string
	// 	want string
	// }{
	// 	{
	// 		name:"hello",
	// 		want:"hello",
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := hello(); got != tt.want {
	// 			t.Errorf("hello() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
