package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "no1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val:   -10,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
		{
			name: "no2",
			args: args{
				nums: []int{1, 3},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		{
			name: "no3",
			args: args{
				nums: []int{3, 5, 8},
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.args.nums)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%q. sortedArrayToBST() = %v, want %v\ndiff = %v", tt.name, got, tt.want, diff)
			}
		})
	}
}
