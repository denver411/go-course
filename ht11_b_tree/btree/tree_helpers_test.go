package btree

import (
	"reflect"
	"testing"
)

func Test_pointsToInts(t *testing.T) {
	type args struct {
		a []*int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"correct", args{[]*int{&k1, &k2, &k3}}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointsToInts(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pointsToInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
