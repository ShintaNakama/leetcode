package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "4",
			args: args{
				nums:   []int{0, 4, 3, 0},
				target: 0,
			},
			want: []int{0, 3},
		},
		{
			name: "5",
			args: args{
				nums:   []int{-1, -2, -3, -4, -5},
				target: -8,
			},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%q. twoSum() = %v, want %v\n diff = %v", tt.name, got, tt.want, diff)
			}
		})
	}
}
