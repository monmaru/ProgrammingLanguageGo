package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

/*
Run:
go build -o ex01
./ex01 1 2 3

Result:
./ex01 1 2 3
*/
