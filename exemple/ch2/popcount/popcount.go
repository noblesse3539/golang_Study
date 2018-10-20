package popCount

import (
	"testing"
)

// pc[i]는 i의 인구수다.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount는 x의 인구수(1로 지정된 비트 수)를 반환한다.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func BenchmarkPopCount(b *testing.B) {
	for i :=0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}
