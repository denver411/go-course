package calc

import (
	"testing"
)

func Test_stackFloat_pop(t *testing.T) {
	s := stackFloat{1, 2, 3}

	tests := []struct {
		name string
		s    *stackFloat
		want float64
	}{
		{"correct work", &s, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.pop(); got != tt.want {
				t.Errorf("stackFloat.pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackFloat_push(t *testing.T) {
	s := stackFloat{1, 2, 3}

	type args struct {
		new float64
	}
	tests := []struct {
		name string
		s    *stackFloat
		args args
		want int
	}{
		{"correct work", &s, args{4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.push(tt.args.new); got != tt.want {
				t.Errorf("stackFloat.push() = %v, want %v", got, tt.want)
			}
		})
	}
}
