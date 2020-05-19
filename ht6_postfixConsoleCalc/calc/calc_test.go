package calc

import (
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		exp string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"correct work", args{"2 + 2 * 2"}, "2 2 2 * +", false},
		{"correct with brackets", args{"(2 + 2) * 2"}, "2 2 + 2 *", false},
		{"error operator in a row", args{"2 + * 2"}, "", true},
		{"error operands in a row", args{"2 + 2 2 * 2"}, "", true},
		{"error unpaired brackets", args{"2 + 2) * 2"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	type args struct {
		exp string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"correct work", args{"2 2 2 * +"}, 6, false},
		{"incorrect work", args{"2 2 2 *"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
