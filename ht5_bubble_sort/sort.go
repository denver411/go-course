package main

import (
	"fmt"
)

// Sort interface
type Sort interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// SortInt struct
type SortInt []int

func (s SortInt) Len() int {
	return len(s)
}

func (s SortInt) Less(i, j int) bool {
	return s[j] > s[i]
}

func (s SortInt) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}

func bubbleSort(s Sort) Sort {
	l := s.Len()
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if s.Less(j, j-1) {
				s.Swap(j-1, j)
			}
		}
	}

	return s
}

func main() {
	// slStr := []string{"ololo", "hey", "hello", "world", "go"}
	sInt := SortInt{6, 4, 2, 9, 1, 0, 6, 2, 8}

	fmt.Println(sInt)
	sortSlInt := bubbleSort(sInt)
	fmt.Println(sortSlInt)
}
