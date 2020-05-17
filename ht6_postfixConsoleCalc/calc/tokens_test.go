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

func Test_token_toString(t *testing.T) {
	tests := []struct {
		name string
		tr   token
		want string
	}{
		{"correct work", tPlus, "+"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.toString(); got != tt.want {
				t.Errorf("token.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createToken(t *testing.T) {
	type args struct {
		t tokenT
		v rune
	}
	tests := []struct {
		name string
		args args
		want token
	}{
		{"correct work", args{operator, '+'}, tPlus},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createToken(tt.args.t, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOpenBracket(t *testing.T) {
	type args struct {
		t token
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"is open bracket", args{tOpenBracket}, true},
		{"is not open bracket", args{tPlus}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOpenBracket(tt.args.t); got != tt.want {
				t.Errorf("isOpenBracket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokensToString(t *testing.T) {
	type args struct {
		ts tokens
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"correct work", args{tokens{tOpenBracket, tOne, tPlus}}, "( 1 +"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokensToString(tt.args.ts); got != tt.want {
				t.Errorf("TokensToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
