package main

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val:   7,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   2,
								Left:  nil,
								Right: nil,
							},
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val:   13,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:  4,
							Left: nil,
							Right: &TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
				targetSum: 22,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
			t.Errorf("%q. hasPathSum() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
