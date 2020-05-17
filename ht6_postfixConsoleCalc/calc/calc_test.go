package calc

import "testing"

func TestCalculate(t *testing.T) {
	ts := tokens{tTwo, tTwo, tPlus, tTwo, tMult}
	wrong := tokens{tTwo, tTwo, tPlus, tTwo}

	type args struct {
		tokens tokens
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"correct work", args{ts}, 8, false},
		{"incorrect work", args{wrong}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.tokens)
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
