package main

import "testing"

func TestDivideWorker(t *testing.T) {
	type args struct {
		inbox chan DivideCommand
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DivideWorker(tt.args.inbox)
		})
	}
}
