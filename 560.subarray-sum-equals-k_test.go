package main

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no1",
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
		{
			name: "no2",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			want: 2,
		},
		{
			name: "no3",
			args: args{
				nums: []int{1, 2, 1, 2, 1},
				k:    3,
			},
			want: 4,
		},
		{
			name: "no4",
			args: args{
				nums: []int{1, -1, 0},
				k:    0,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subarraySum(tt.args.nums, tt.args.k)
			if got != tt.want {
				t.Errorf("%q. subarraySum() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
