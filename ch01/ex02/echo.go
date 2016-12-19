package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args[:] {
		fmt.Println(strconv.Itoa(i) + " " + arg)
	}
}

/*
Run:
go build -o ex01
./ex01 a b c

Result:
0 ./ex02
1 a
2 b
3 c
*/
