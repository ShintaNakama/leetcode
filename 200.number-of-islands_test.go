package main

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no1",
			args: args{
				grid: [][]byte{
					{49, 49, 49, 49, 48},
					{49, 49, 48, 49, 48},
					{49, 49, 48, 48, 48},
					{48, 48, 48, 48, 48},
				},
			},
			want: 1,
		},
		{
			name: "no2",
			args: args{
				grid: [][]byte{
					{49, 49, 48, 48, 48},
					{49, 49, 48, 48, 48},
					{48, 48, 49, 48, 48},
					{48, 48, 48, 49, 49},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("%q. numIslands() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
