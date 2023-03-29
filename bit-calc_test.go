package main

import "testing"

func Test_andCalc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "no.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			andCalc()
		})
	}
}
