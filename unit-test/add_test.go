package main

import (
	"reflect"
	"testing"
)

type args[T int|float32|float64] struct {
	n1 T
	n2 T
}

type testCase[T int|float32|float64] struct{
	name string
	args args[T]
	want T
}

func testAdd[T int|float32|float64](t *testing.T, testCases []testCase[T]) {  
	for _, tt := range testCases {  
		t.Run(tt.name, func(t *testing.T) {  
			if got := Add(tt.args.n1, tt.args.n2); !reflect.DeepEqual(got, tt.want) {  
				t.Errorf("Add() = %v, want %v", got, tt.want)  
			}  
		})  
	}  
}  
func TestAdd(t *testing.T) {

	testint := []testCase[int] {
		{
			name: "AddInt",  
			args: args[int]{n1: 2, n2: 3},  
			want: 5, 
		},
		{
			name: "AddInt",  
			args: args[int]{n1: 1000, n2: 3},  
			want: 1003, 
		},
	}
	testAdd(t, testint)
	

	testfloat32 := []testCase[float32] {
		{
			name: "AddInt",  
			args: args[float32]{n1: 2.0, n2: -0.2},  
			want: 1.8, 
		},
		{
			name: "AddInt",  
			args: args[float32]{n1: 10.81, n2: 3},  
			want: 13.81, 
		},
	}

	testAdd(t, testfloat32)

}
