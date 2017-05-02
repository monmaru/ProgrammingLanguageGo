package popcount4

import (
	"testing"

	p2 "github.com/monmaru/ProgrammingLanguageGo/ch02/ex03/popcount2"
	p3 "github.com/monmaru/ProgrammingLanguageGo/ch02/ex04/popcount3"
)

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p2.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p3.PopCount3(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

/*
go test -bench=.
BenchmarkPopCount2-4    100000000               21.2 ns/op
BenchmarkPopCount3-4    20000000                94.4 ns/op
BenchmarkPopCount4-4    50000000                22.4 ns/op
PASS
ok      github.com/monmaru/ProgrammingLanguageGo/ch02/ex05/popcount4    5.345s
*/
