package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	reverse(a)
	fmt.Println(a)
	
	// n 만큼 좌측으로 회전의 경우 reverse(a[:n]) reverse(a[n:]) reverse(a[:]) 하면 됨
	reverse(a)
	// 3만큼 좌측으로 회전
	reverse(a[:3]) 
	reverse(a[3:]) 
	reverse(a)
	fmt.Println(a)
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
