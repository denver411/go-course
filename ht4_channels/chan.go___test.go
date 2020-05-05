package main

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		ch <-chan int
	}
	ch := make(chan int, 1000)
	ch <- 2
	ch <- 4
	ch <- 22
	close(ch)
	ch2 := make(chan int, 1000)
	close(ch2)
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ok", args{ch}, 28},
		{"notOk", args{ch2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.ch); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
