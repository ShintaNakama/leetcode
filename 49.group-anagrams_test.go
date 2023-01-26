package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "no1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%q. groupAnagrams() = %v, want %v\n diff = %v", tt.name, got, tt.want, diff)
			}
		})
	}
}
