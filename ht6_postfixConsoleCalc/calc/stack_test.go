package calc

import (
	"reflect"
	"testing"
)

func Test_tokens_last(t *testing.T) {
	ts := tokens{tOne, tPlus, tTwo}

	tests := []struct {
		name string
		s    tokens
		want token
	}{
		{"correct work", ts, tTwo},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.last(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokens.last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokens_pop(t *testing.T) {
	ts := tokens{tOne, tPlus, tTwo}

	tests := []struct {
		name string
		s    *tokens
		want token
	}{
		{"correct work", &ts, tTwo},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokens.pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokens_push(t *testing.T) {
	ts := tokens{tOne, tPlus, tTwo}

	type args struct {
		new token
	}
	tests := []struct {
		name string
		s    *tokens
		args args
		want int
	}{
		{"correct work", &ts, args{tThree}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.push(tt.args.new); got != tt.want {
				t.Errorf("tokens.push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokens_empty(t *testing.T) {
	ts := tokens{tOne, tPlus, tTwo}

	tests := []struct {
		name string
		s    tokens
		want bool
	}{
		{"is not empty", ts, false},
		{"is empty", tokens{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.empty(); got != tt.want {
				t.Errorf("tokens.empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
