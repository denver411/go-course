package main

import (
	"reflect"
	"testing"
)

func TestSortIntLen(t *testing.T) {
	var tests = []struct {
		name string
		data []int
		want int
	}{
		{"ok", []int{6, 4, 2, 9, 1, 0, 6, 2, 8}, 9},
		{"emptyArray", []int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := SortInt{data: tt.data}
			got := c.Len()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divide() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestSortIntLen(t *testing.T) {
// 	type args struct {
// 		a interface{}
// 		b interface{}
// 	}
// 	var tests = []struct {
// 		name    string
// 		args    args
// 		want    interface{}
// 		wantErr bool
// 	}{
// 		{"ok", args{6., 4.}, 1.5, false},
// 		{"divByZeroErr", args{6., 0.}, 0, true},
// 		{"argTypeErr", args{"six", 0.}, 0, true},
// 		{"noArgErr", args{}, 0, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := Calc{}
// 			got, err := c.Divide(tt.args.a, tt.args.b)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Divide() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
