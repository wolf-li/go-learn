package main

import (
	"testing"
	"github.com/gin-gonic/gin"
)

func Test_queryform(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryform(tt.args.ctx)
		})
	}
}
