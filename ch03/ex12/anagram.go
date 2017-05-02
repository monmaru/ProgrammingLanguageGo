package ex12

import "reflect"

func isAnagram(s1, s2 string) bool {
	return reflect.DeepEqual(countRunes(s1), countRunes(s2))
}

func countRunes(s string) map[rune]int {
	c := map[rune]int{}
	for _, r := range s {
		c[r]++
	}
	return c
}
