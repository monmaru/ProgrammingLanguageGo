package main

import (
	"crypto/sha256"
	"fmt"

	p "github.com/monmaru/ProgrammingLanguageGo/ch02/ex03/popcount2"
)

const hashSize = sha256.Size

func main() {
	var v1, v2 string
	fmt.Print("input text v1 > ")
	fmt.Scan(&v1)
	fmt.Print("input text v2 > ")
	fmt.Scan(&v2)

	d1 := sha256.Sum256([]byte(v1))
	d2 := sha256.Sum256([]byte(v2))

	fmt.Printf("popDiffCount(%v, %v): %v\n", v1, v2, popDiffCount(d1, d2))
}

func popDiffCount(d1, d2 [hashSize]byte) (count int) {
	for i := 0; i < hashSize; i++ {
		count += p.PopCount(uint64(d1[i]) ^ uint64(d2[i]))
	}
	return
}
