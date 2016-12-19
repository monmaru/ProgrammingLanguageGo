package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("Hit count is %d.\t%s in [%s]\n", count, line, strings.Join(fileNames[line], ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileNames[input.Text()] = append(fileNames[input.Text()], f.Name())
	}
}
