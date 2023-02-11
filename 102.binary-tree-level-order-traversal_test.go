package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "no1",
			args: args{
				root: &TreeNode{
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
			want: [][]int{
				{3}, {9, 20}, {15, 7},
			},
		},
		{
			name: "no2",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: [][]int{{1}},
		},
		{
			name: "no3",
			args: args{
				root: nil,
			},
			want: [][]int{},
		},
		{
			name: "no4([0,2,4,1,null,3,-1,5,1,null,6,null,8])",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
							Left: &TreeNode{
								Val:   5,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:  3,
							Left: nil,
							Right: &TreeNode{
								Val:   6,
								Left:  nil,
								Right: nil,
							},
						},
						Right: &TreeNode{
							Val:  -1,
							Left: nil,
							Right: &TreeNode{
								Val:   8,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
			},
			want: [][]int{
				{0}, {2, 4}, {1, 3, -1}, {5, 1, 6, 8},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.args.root)
			//got := levelOrder2(tt.args.root)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%q. levelOrder() = %v, want %v\n diff = %v", tt.name, got, tt.want, diff)
			}
		})
	}
}
