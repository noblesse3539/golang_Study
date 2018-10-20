package main

import (
	"fmt"
	
)

func main() {
	var a []int
	
	fmt.Println(len(a))
	fmt.Println(cap(a))
	
	a = make([]int, 1,10)
	
	b := a[:4]
	
	fmt.Println(len(b))
	fmt.Println(cap(b))
}