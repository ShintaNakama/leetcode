package main

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no1",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "no2",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "no3",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
		{
			name: "no4",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 4,
			},
			want: 2,
		},
		{
			name: "no5",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			want: 0,
		},
		{
			name: "no6",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: 0,
		},
		{
			name: "no7",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "no8",
			args: args{
				nums:   []int{1},
				target: 2,
			},
			want: 1,
		},
		{
			name: "no9",
			args: args{
				nums:   []int{1, 3, 5},
				target: 6,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("%q. searchInsert() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
