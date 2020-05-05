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
type SortInt struct {
	data []int
}

func (s SortInt) Len() int {
	return len(s.data)
}

func (s SortInt) Less(i, j int) bool {
	return s.data[j] > s.data[i]
}

func (s SortInt) Swap(i, j int) {
	s.data[j], s.data[i] = s.data[i], s.data[j]
}

func bubbleSort(s Sort) Sort {
	l := s.Len()
	for i := 0; i < l; i++ {
		for j := 1; j < l; j++ {
			if s.Less(j, j-1) {
				s.Swap(j-1, j)
			}
		}
	}

	return s
}

func main() {
	// slStr := []string{"ololo", "hey", "hello", "world", "go"}
	sInt := SortInt{
		data: []int{6, 4, 2, 9, 1, 0, 6, 2, 8},
	}

	fmt.Println(sInt)
	sortSlInt := bubbleSort(sInt)
	fmt.Println(sortSlInt)
}
