package calc

import (
	"reflect"
	"testing"
)

func Test_handleBracket(t *testing.T) {
	stack := &tokens{tOpenBracket, tPlus, tTwo}
	stackNoBrackets := &tokens{tPlus, tTwo}
	res := &tokens{}

	type args struct {
		stack *tokens
		res   *tokens
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"correct work", args{stack, res}, false},
		{"unpaired brackets", args{stackNoBrackets, res}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handleBracket(tt.args.stack, tt.args.res); (err != nil) != tt.wantErr {
				t.Errorf("handleBracket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_makePostfix(t *testing.T) {
	tokensInfixSimple := tokens{tOne, tPlus, tTwo}
	tokensPostfixSimple := tokens{tOne, tTwo, tPlus}
	tokensInfixPrioritized := tokens{tTwo, tPlus, tTwo, tMult, tTwo}
	tokensPostfixPrioritized := tokens{tTwo, tTwo, tTwo, tMult, tPlus}
	tokensInfixBrackets := tokens{tOpenBracket, tTwo, tPlus, tTwo, tCloseBracket, tMult, tTwo}
	tokensPostfixBrackets := tokens{tTwo, tTwo, tPlus, tTwo, tMult}

	tokensInfixOperatorsInRow := tokens{tTwo, tPlus, tPlus, tTwo}
	tokensInfixLiteralsInRow := tokens{tTwo, tTwo, tPlus}
	tokensInfixBracketsError := tokens{tTwo, tPlus, tTwo, tCloseBracket}

	type args struct {
		infix tokens
	}
	tests := []struct {
		name    string
		args    args
		want    tokens
		wantErr bool
	}{
		{"simple expression", args{tokensInfixSimple}, tokensPostfixSimple, false},
		{"prioritized expression", args{tokensInfixPrioritized}, tokensPostfixPrioritized, false},
		{"expression with brackets", args{tokensInfixBrackets}, tokensPostfixBrackets, false},
		{"two operators in row error", args{tokensInfixOperatorsInRow}, nil, true},
		{"two literals in row error", args{tokensInfixLiteralsInRow}, nil, true},
		{"brackets error", args{tokensInfixBracketsError}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makePostfix(tt.args.infix)
			if (err != nil) != tt.wantErr {
				t.Errorf("makePostfix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makePostfix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleOperator(t *testing.T) {
	type args struct {
		stack   *tokens
		res     *tokens
		current token
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleOperator(tt.args.stack, tt.args.res, tt.args.current)
		})
	}
}
