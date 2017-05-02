package ex11

import (
	"bytes"
	"strings"
)

const base = 3

func comma(s string) string {
	integer, fractional := separate(s)
	buf := bytes.Buffer{}
	writeInteger(&buf, integer)
	writeFractional(&buf, fractional)
	return buf.String()
}

func separate(s string) (integer, fractional string) {
	sp := strings.Split(s, ".")
	integer = sp[0]
	if len(sp) > 1 {
		fractional = sp[1]
	}
	return
}

func writeInteger(buf *bytes.Buffer, integer string) {
	if len(integer) > 0 && (integer[0] == '+' || integer[0] == '-') {
		buf.WriteRune(rune(integer[0]))
		integer = integer[1:]
	}

	n := len(integer)
	m := n % base
	if m == 0 && n >= base {
		m = base
	}
	buf.WriteString(integer[:m])
	integer = integer[m:]

	for len(integer) > 0 {
		buf.WriteString(",")
		buf.WriteString(integer[:base])
		integer = integer[base:]
	}
}

func writeFractional(buf *bytes.Buffer, fractional string) {
	if len(fractional) > 0 {
		buf.WriteString(".")
		buf.WriteString(fractional)
	}
}
