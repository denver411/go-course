package main

import (
	"reflect"
	"testing"
)

func TestSortIntLen(t *testing.T) {
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

func TestSortIntLess(t *testing.T) {
	sInt := SortInt{6, 4, 2, 9, 1, 0, 6, 2, 8}
	var tests = []struct {
		name string
		i    int
		j    int
		data SortInt
		want bool
	}{
		{"i less than j", 2, 3, sInt, true},
		{"i more than j", 0, 1, sInt, false},
		{"i equal j", 2, 7, sInt, false},
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

func TestSortIntSwap(t *testing.T) {
	sInt := SortInt{6, 4, 2, 9, 1, 0, 6, 2, 8}
	var tests = []struct {
		name string
		i    int
		j    int
		data SortInt
		want bool
	}{
		{"elements 0,1 swap", 0, 1, sInt, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prevI, prevJ := tt.data[tt.i], tt.data[tt.j]
			tt.data.Swap(tt.i, tt.j)
			got := prevI == tt.data[tt.j] && prevJ == tt.data[tt.i]

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
