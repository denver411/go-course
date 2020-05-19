package calc

type stackFloat []float64

func (s *stackFloat) pop() float64 {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func (s *stackFloat) push(new float64) int {
	*s = append(*s, new)
	return len(*s)
}
