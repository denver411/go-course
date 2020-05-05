package main

import (
	"reflect"
	"testing"
)

func TestCalc_Divide(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	var tests = []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"ok", args{6., 4.}, 1.5, false},
		{"divByZeroErr", args{6., 0.}, 0, true},
		{"argTypeErr", args{"six", 0.}, 0, true},
		{"noArgErr", args{}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Calc{}
			got, err := c.Divide(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divide() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalc_Multiply(t *testing.T) {
	type a []interface{}
	type args struct {
		a a
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"ok", args{a{6, 4}}, 24, false},
		{"argTypeErr", args{a{"six", 4}}, 0, true},
		{"noArgErr", args{a{}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Calc{}
			got, err := c.Multiply(tt.args.a...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalc_Substruct(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"ok", args{6, 4}, 2, false},
		{"argTypeErr", args{"six", 22}, 0, true},
		{"noArgErr", args{}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Calc{}
			got, err := c.Substruct(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Substruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Substruct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalc_Sum(t *testing.T) {
	type a []interface{}
	type args struct {
		a a
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"ok", args{a{6, 5}}, 11, false},
		{"argTypeErr", args{a{"six", 4}}, 0, true},
		{"noArgErr", args{a{}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Calc{}
			got, err := c.Sum(tt.args.a...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() got = %v, want %v", got, tt.want)
			}
		})
	}
}
