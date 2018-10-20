package popcount_test

import (
	"testing"
	"exercise/ch2/exer2_3/popcount"
)
func PopCount1(x uint64) int {
	result := 0
	for i := x; i != 0; i /= 2 {
		if i&1 != 0 {
			result++
		}
	}
	return result
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}
func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(0x1234567890ABCDEF)
	}
}
