package ex04

func rotate(s []int) {
	first := s[0]
	copy(s, s[1:])
	s[len(s)-1] = first
}
