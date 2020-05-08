package calc

// Stack interface
type stackInt interface {
	Len() int
	Pop() string
	Push(string) int
}

// Stack struct
type Stack []string

// Len method
func (s Stack) Len() int {
	return len(s)
}

// Last method
func (s Stack) Last() string {
	return s[s.Len()-1]
}

// Pop method
func (s *Stack) Pop() string {
	st := *s
	last := st.Last()
	*s = st[:st.Len()-1]
	return last
}

// Push method
func (s *Stack) Push(new string) int {
	*s = append(*s, new)
	return s.Len()
}
