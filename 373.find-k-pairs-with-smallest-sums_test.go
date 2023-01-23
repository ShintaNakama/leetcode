package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "no1",
			args: args{
				nums1: []int{1, 7, 11},
				nums2: []int{2, 4, 6},
				k:     3,
			},
			want: [][]int{
				{1, 2},
				{1, 4},
				{1, 6},
			},
		},
		{
			name: "no2",
			args: args{
				nums1: []int{1, 1, 2},
				nums2: []int{1, 2, 3},
				k:     2,
			},
			want: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		{
			name: "no3",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3},
				k:     3,
			},
			want: [][]int{
				{1, 3},
				{2, 3},
			},
		},
		{
			name: "no4",
			args: args{
				nums1: []int{1, 1, 2},
				nums2: []int{1, 2, 3},
				k:     10,
			},
			want: [][]int{
				{1, 1},
				{1, 1},
				{1, 2},
				{2, 1},
				{1, 2},
				{2, 1},
				{2, 2},
				{1, 3},
				{1, 3},
				{2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%q. kSmallestPairs() = %v, want %v,\ndiff = %v", tt.name, got, tt.want, diff)
			}
		})
	}
}
