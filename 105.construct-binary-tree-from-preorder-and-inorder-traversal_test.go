package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "no1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			name: "no2",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: &TreeNode{
				Val: -1,
			},
		},
		{
			name: "no3",
			args: args{
				preorder: []int{3, 9, 4, 8, 20, 15, 7},
				inorder:  []int{4, 9, 8, 3, 15, 20, 7},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   8,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.args.preorder, tt.args.inorder)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%q. buildTree() = %v, want %v\n diff = %v", tt.name, got, tt.want, diff)
			}
		})
	}
}
