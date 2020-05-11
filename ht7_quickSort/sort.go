package main

import (
	"fmt"
)

// Sort interface
type Sort interface {
	Len() int
	Less(i, j int) bool
}

// SortInt struct
type SortInt []int

func (s SortInt) Len() int {
	return len(s)
}

func (s SortInt) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortInt) Swap(i, j int) {
	if i != j {
		s[j], s[i] = s[i], s[j]
	}
}

func partition(s SortInt, first int, last int) int {
	if first >= last {
		return 0
	}
	x := first
	for j := first + 1; j <= last; j++ {
		if s.Less(j, first) {
			x = x + 1
			s.Swap(x, j)
		}
	}
	s.Swap(x, first)
	return x
}

func sort(s SortInt, first int, last int) {
	if first < 0 || last > s.Len()-1 {
		return
	}
	if first < last {
		i := partition(s, first, last)
		sort(s, first, i-1)
		sort(s, i+1, last)
	}
}

func quickSort(s SortInt) SortInt {
	if s.Len() <= 1 {
		return s
	}

	sort(s, 0, s.Len()-1)

	return s
}

func main() {
	sInt := SortInt{6, 4, 2, 9, 1, 0, 6, 2, 8}
	fmt.Println("before:", sInt)
	sortSlInt := quickSort(sInt)
	fmt.Println("after:", sortSlInt)
}
