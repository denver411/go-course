package calc

import "testing"

func Test_applyOperator(t *testing.T) {
	type args struct {
		operator string
		operand1 float64
		operand2 float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"sum", args{"+", 2, 2}, 4, false},
		{"mul", args{"*", 2, 2}, 4, false},
		{"div", args{"/", 2, 2}, 1, false},
		{"div by zero", args{"/", 2, 0}, 0, true},
		{"sub", args{"-", 2, 2}, 0, false},
		{"pow", args{"^", 2, 2}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applyOperator(tt.args.operator, tt.args.operand1, tt.args.operand2)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyOperator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("applyOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLessPriority(t *testing.T) {
	type args struct {
		operator1 string
		operator2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"less priority", args{"+", "+"}, true},
		{"more priority", args{"*", "+"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLessPriority(tt.args.operator1, tt.args.operator2); got != tt.want {
				t.Errorf("isLessPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
