package main

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no1",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "no2",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
			},
			want: 12,
		},
		{
			name: "no3",
			args: args{
				nums: []int{2, 1, 1, 2},
			},
			want: 4,
		},
		{
			name: "no4",
			args: args{
				nums: []int{1, 2, 1, 1},
			},
			want: 3,
		},
		{
			name: "no5",
			args: args{
				nums: []int{6, 3, 10, 8, 2, 10, 3, 5, 10, 5, 3},
			},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("%q. rob() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
