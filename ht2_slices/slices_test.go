package main

import (
	"reflect"
	"testing"
)

func TestSliceReverse(t *testing.T) {
	type sl []interface{}
	type args struct {
		sl sl
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"okInt", args{sl{1, 2, 3, 4, 5, 6, 7}}, sl{7, 6, 5, 4, 3, 2, 1}},
		{"okString", args{sl{"one", "two", "three"}}, sl{"three", "two", "one"}},
		{"okOneArg", args{sl{1}}, sl{1}},
		{"okNoArg", args{sl{}}, sl{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceReverse(tt.args.sl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func
TestSliceWithStep(t *testing.T) {
	type sl []interface{}
	type args struct {
		sl   []interface{}
		step int64
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"okShift3", args{sl{1, 2, 3, 4, 5, 6, 7}, 3}, sl{5, 6, 7, 1, 2, 3, 4}},
		{"okShift0", args{sl{1, 2, 3, 4, 5, 6, 7}, 0}, sl{1, 2, 3, 4, 5, 6, 7}},
		{"okStringShift1", args{sl{"one", "two", "three"}, 1}, sl{"three", "one", "two"}},
		{"okOneArg", args{sl{1},5}, sl{1}},
		{"okNoArg", args{sl{}, 3}, sl{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceWithStep(tt.args.sl, tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceWithStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
