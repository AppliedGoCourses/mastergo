package main

import "testing"

func Test_collatz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2", args{2}, 1},
		{"3", args{3}, 7},
		{"4", args{4}, 2},
		{"5", args{5}, 5},
		{"6", args{6}, 8},
		{"7", args{7}, 16},
		{"999", args{999}, 49},
		{"1000", args{1000}, 111},
		{"1001", args{1001}, 142},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collatz(tt.args.n); got != tt.want {
				t.Errorf("collatz() = %v, want %v", got, tt.want)
			}
		})
	}
}
