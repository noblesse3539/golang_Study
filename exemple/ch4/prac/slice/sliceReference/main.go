package main

import (
	"fmt"
)

func main() {
	intSlice := []int{0,1,2,3,4,5}
	intArray := [...]int{0,1,2,3,4,5}
	fmt.Printf("슬라이스	길이:%d 용량:%d\n", len(intSlice), cap(intSlice))
	fmt.Printf("배열		길이:%d 용량:%d\n", len(intArray), cap(intArray))
	
	tempSlice := intArray[:]
//	a(intSlice)
	a(&intArray)
	b(intSlice)
	b(&intArray)
	fmt.Printf("%p	%p\n", &intArray, tempSlice)
	fmt.Printf("%d	%d\n", intSlice[0], intArray[0])
}

func a(x *[6]int) {
	x[0] = 10
}
func b(x []int) {
	x[0] = 20
}