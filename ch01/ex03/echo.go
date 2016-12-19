package main

import "strings"

func echo02(input []string) {
	var s, sep string
	for _, arg := range input {
		s += sep + arg
		sep = " "
	}
}

func echo03(input []string) {
	strings.Join(input, " ")
}
