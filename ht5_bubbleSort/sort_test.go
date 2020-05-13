package main

import (
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	var tests = []struct {
		name string
		data SortInt
		want int
	}{
		{"ok", SortInt{6, 4, 2, 9, 1, 0, 6, 2, 8}, 9},
		{"emptyArray", SortInt{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.data.Len()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Len() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLess(t *testing.T) {
	sInt := SortInt{6, 4, 2, 6}
	var tests = []struct {
		name string
		i    int
		j    int
		data SortInt
		want bool
	}{
		{"i less than j", 2, 3, sInt, true},
		{"i more than j", 0, 1, sInt, false},
		{"i equal j", 0, 3, sInt, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.data.Less(tt.i, tt.j)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Less() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	s := SortInt{6, 4}
	sSwap := SortInt{4, 6}
	var tests = []struct {
		name string
		i    int
		j    int
		data SortInt
		want SortInt
	}{
		{"elements swapped", 0, 1, s, sSwap},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.data.Swap(tt.i, tt.j)
			got := tt.data

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Swap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	sIntUnsort := SortInt{6, 4, 2, 9, 1, 0, 6, 2, 8}
	sIntSort := SortInt{0, 1, 2, 2, 4, 6, 6, 8, 9}
	sIntEmpty := SortInt{}

	var tests = []struct {
		name string
		data SortInt
		want SortInt
	}{
		{"is sorted", sIntUnsort, sIntSort},
		{"empty", sIntEmpty, sIntEmpty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bubbleSort(tt.data)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() got = %v, want %v", got, tt.want)
			}
		})
	}
}
