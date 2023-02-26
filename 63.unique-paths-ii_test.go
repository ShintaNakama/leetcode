package main

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no1",
			args: args{
				obstacleGrid: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: 2,
		},
		{
			name: "no2",
			args: args{
				obstacleGrid: [][]int{
					{0, 1},
					{0, 0},
				},
			},
			want: 1,
		},
		{
			name: "no3",
			args: args{
				obstacleGrid: [][]int{
					{1, 0},
				},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("%q. uniquePathsWithObstacles() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
