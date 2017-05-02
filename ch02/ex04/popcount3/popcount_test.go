package popcount3

import (
	"testing"

	p "github.com/monmaru/ProgrammingLanguageGo/ch02/ex03/popcount2"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(0x1234567890ABCDEF)
	}
}

/*
go test -bench=.
BenchmarkPopCount-4     100000000               20.9 ns/op
BenchmarkPopCount3-4    20000000                93.6 ns/op
PASS
ok      github.com/monmaru/ProgrammingLanguageGo/ch02/ex04/popcount3    4.146s
*/
