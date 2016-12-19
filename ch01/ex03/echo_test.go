package main

import (
	"fmt"
	"testing"
	"time"
)

var input = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

func BenchmarkEcho02(b *testing.B) {
	for i := 1; i < b.N; i++ {
		echo02(input)
	}
}

func BenchmarkEcho03(b *testing.B) {
	for i := 1; i < b.N; i++ {
		echo03(input)
	}
}

func stopWatch(loopCount int, fn func()) {
	start := time.Now()
	for i := 0; i < loopCount; i++ {
		fn()
	}
	fmt.Printf("Execution time: %.2fs\n", time.Since(start).Seconds())
}

func TestEcho(t *testing.T) {
	const loopCount = 10000000
	stopWatch(loopCount, func() { echo02(input) })
	stopWatch(loopCount, func() { echo03(input) })
}

/*
Run:
go test -bench .

Result:
Execution time: 3.97s
Execution time: 1.45s
BenchmarkEcho02-4        3000000               402 ns/op
BenchmarkEcho03-4       10000000               141 ns/op
*/
