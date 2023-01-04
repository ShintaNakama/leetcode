package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want KthLargest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Constructor(tt.args.k, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Constructor() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestKthLargest_Add(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *KthLargest
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		this := &KthLargest{}
		if got := this.Add(tt.args.val); got != tt.want {
			t.Errorf("%q. KthLargest.Add() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
