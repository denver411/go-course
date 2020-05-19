package calc

import (
	"reflect"
	"testing"
)

var tPlus = token{t: operator, v: "+"}
var tMinus = token{t: operator, v: "-"}
var tMult = token{t: operator, v: "*"}
var tOne = token{t: literal, v: "1"}
var tTwo = token{t: literal, v: "2"}
var tThree = token{t: literal, v: "3"}
var tOpenBracket = token{t: bracket, v: "("}
var tCloseBracket = token{t: bracket, v: ")"}

func Test_createToken(t *testing.T) {
	type args struct {
		t tokenT
		v string
	}
	tests := []struct {
		name string
		args args
		want token
	}{
		{"correct work", args{operator, "+"}, tPlus},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createToken(tt.args.t, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_token_String(t *testing.T) {
	tests := []struct {
		name string
		tr   token
		want string
	}{
		{"correct work", tPlus, "+"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("token.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_token_isOpenBracket(t *testing.T) {
	tests := []struct {
		name string
		tr   token
		want bool
	}{
		{"is open bracket", tOpenBracket, true},
		{"is not open bracket", tPlus, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.isOpenBracket(); got != tt.want {
				t.Errorf("token.isOpenBracket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokens_String(t *testing.T) {
	tests := []struct {
		name string
		ts   tokens
		want string
	}{
		{"correct work", tokens{tOpenBracket, tOne, tPlus}, "( 1 +"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.String(); got != tt.want {
				t.Errorf("tokens.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokens_last(t *testing.T) {
	ts := tokens{tOne, tPlus, tTwo}

	tests := []struct {
		name string
		ts   tokens
		want token
	}{
		{"correct work", ts, tTwo},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.last(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokens.last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokens_pop(t *testing.T) {
	ts := tokens{tOne, tPlus, tTwo}

	tests := []struct {
		name string
		ts   *tokens
		want token
	}{
		{"correct work", &ts, tTwo},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.pop(); !reflect.DeepEqual(got, tt.want) {
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
		ts   *tokens
		args args
		want int
	}{
		{"correct work", &ts, args{tThree}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.push(tt.args.new); got != tt.want {
				t.Errorf("tokens.push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokens_empty(t *testing.T) {
	ts := tokens{tOne, tPlus, tTwo}

	tests := []struct {
		name string
		ts   tokens
		want bool
	}{
		{"is not empty", ts, false},
		{"is empty", tokens{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.empty(); got != tt.want {
				t.Errorf("tokens.empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
