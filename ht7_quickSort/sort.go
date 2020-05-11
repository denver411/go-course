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
	return s[j] > s[i]
}

func quickSortRec(s SortInt) SortInt {
	if s.Len() <= 1 {
		return s
	}

	x := s[0]
	xs := s[1:]
	less := SortInt{}
	more := SortInt{}

	for idx, i := range xs {
		if s.Less(idx+1, 0) {
			less = append(less, i)
		} else {
			more = append(more, i)
		}
	}

	left := quickSortRec(less)
	right := quickSortRec(more)
	right = append(SortInt{x}, right...)

	return append(left, right...)
}

func main() {
	sInt := SortInt{6, 4, 2, 9, 1, 0, 6, 2, 8}

	fmt.Println("before:", sInt)
	sortSlInt := quickSortRec(sInt)
	fmt.Println("after:", sortSlInt)
}
