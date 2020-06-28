package calc

import (
	"reflect"
	"testing"
)

func Test_tokenize(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		want    tokens
		wantErr bool
	}{
		{"correct work", args{"1 + 2 - 3"}, tokens{tOne, tPlus, tTwo, tMinus, tThree}, false},
		{"wrong symbol", args{"1 + y - 3"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tokenize(tt.args.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
