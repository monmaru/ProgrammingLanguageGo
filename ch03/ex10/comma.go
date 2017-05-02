package ex10

import "bytes"

const base = 3

func comma(s string) string {
	if len(s) <= base {
		return s
	}

	buf := bytes.Buffer{}
	n := len(s)
	buf.Grow(n + n/base)

	m := n % base
	if m == 0 {
		m = base
	}

	buf.WriteString(s[:m])
	s = s[m:]

	for len(s) > 0 {
		buf.WriteString(",")
		buf.WriteString(s[:base])
		s = s[base:]
	}
	return buf.String()
}
