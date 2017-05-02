package ex12

import "testing"

var tests = []struct {
	in1, in2 string
	want     bool
}{
	{"", "", true},
	{"aA", "Aa", true},
	{"abc", "cba", true},
	{"ほげもげ", "もほげげ", true},
	{"a", "z", false},
	{"aaA", "aAA", false},
	{"ほげー", "もげー", false},
}

func TestComma(t *testing.T) {
	for _, tt := range tests {
		if ret := isAnagram(tt.in1, tt.in2); ret != tt.want {
			t.Errorf("isAnagram(%q, %q) = %v; want %v", tt.in1, tt.in2, ret, tt.want)
		}
	}
}
