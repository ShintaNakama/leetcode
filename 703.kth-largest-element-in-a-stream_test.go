package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
		{
			name: "no1",
			args: args{
				k:    3,
				nums: []int{4, 5, 8, 2},
			},
			want: KthLargest{
				K:    3,
				Nums: &IntHeap{4, 5, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Constructor(tt.args.k, tt.args.nums)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%q. Constructor() = %v, want %v\ndiff = %v", tt.name, got, tt.want, diff)
			}
		})
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
		{
			name: "no1",
			this: &KthLargest{
				K:    3,
				Nums: &IntHeap{4, 5, 8},
			},
			args: args{
				val: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.this
			got := this.Add(tt.args.val)
			t.Log(this.Nums)
			if got != tt.want {
				t.Errorf("%q. KthLargest.Add() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
