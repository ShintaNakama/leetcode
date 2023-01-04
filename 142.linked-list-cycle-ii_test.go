package main

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. detectCycle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
