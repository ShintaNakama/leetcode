package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
		{
			name: "no2",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			want: 0,
		},
		{
			name: "no3",
			args: args{
				beginWord: "a",
				endWord:   "c",
				wordList:  []string{"a", "b", "c"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%q. ladderLength() = %v, want %v\ndiff = %v", tt.name, got, tt.want, diff)
			}
		})
	}
}
