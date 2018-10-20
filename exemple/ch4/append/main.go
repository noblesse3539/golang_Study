package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int { // 반복문 사용
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]

	} else {
		zcap := zlen // len(x) 가 0일 경우 zcap을 1로 늘릴 수 있어.
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		for i, n := range x {
			z[i] = n
		}
	}
	z[len(x)] = y
	return z
}

func appendInt2(x []int, y int) []int { // 내장함수 copy사용.
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]

	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
func appendInt3(x []int, y ...int) []int { // 가변인자.
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]

	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
func main() {
	var y []int
	fmt.Printf("%v %d %d\n", y, len(y), cap(y))
	y = appendInt(y, 10)
	fmt.Printf("%v %d %d\n", y, len(y), cap(y))
	y = appendInt(y, 20)
	fmt.Printf("%v %d %d\n", y, len(y), cap(y))
	y = appendInt(y, 30)
	fmt.Printf("%v %d %d\n", y, len(y), cap(y))
	y = appendInt(y, 40)
	fmt.Printf("%v %d %d\n", y, len(y), cap(y))
	y = appendInt(y, 50)
	fmt.Printf("%v %d %d\n", y, len(y), cap(y))

	fmt.Println(">>><<<")
	var x []int
	fmt.Printf("%v %d %d\n", x, len(x), cap(x))
	x = appendInt2(x, 10)
	fmt.Printf("%v %d %d\n", x, len(x), cap(x))
	x = appendInt2(x, 20)
	fmt.Printf("%v %d %d\n", x, len(x), cap(x))
	x = appendInt2(x, 30)
	fmt.Printf("%v %d %d\n", x, len(x), cap(x))
	x = appendInt2(x, 40)
	fmt.Printf("%v %d %d\n", x, len(x), cap(x))
	x = appendInt2(x, 50)
	fmt.Printf("%v %d %d\n", x, len(x), cap(x))
}
