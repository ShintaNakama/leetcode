package main

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no1(2,3,7,101)",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "no2",
			args: args{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "no3",
			args: args{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
		{
			name: "no4",
			args: args{
				nums: []int{4, 10, 4, 3, 8, 9},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLIS(tt.args.nums)
			if got != tt.want {
				t.Errorf("%q. lengthOfLIS() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
