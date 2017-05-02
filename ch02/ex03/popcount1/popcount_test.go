package popcount1

import "testing"

const size = 256

func BenchmarkPopCount(b *testing.B) {
	for i := uint64(0); i < size; i++ {
		PopCount(i)
	}
}
