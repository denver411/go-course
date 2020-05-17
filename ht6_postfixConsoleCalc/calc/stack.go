package calc

func (s tokens) last() token {
	return s[len(s)-1]
}

func (s *tokens) pop() token {
	last := s.last()
	*s = (*s)[:len(*s)-1]
	return last
}

func (s *tokens) push(new token) int {
	*s = append(*s, new)
	return len(*s)
}

func (s tokens) empty() bool {
	return len(s) == 0
}
