package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))
	
	c := [32]byte{0:0xff, 2: 0xea}
	d := [32]byte{2:0xff}
	fmt.Println(BitCount(a,b))
	fmt.Println(BitCount(c,d))
	fmt.Printf("%x\n", a)
	fmt.Printf("%x\n", c)
}

func BitCount(a, b [32]byte) int {
	c := [32]byte{}
	n := 0
	for i := 0; i< len(a); i++ {
		c[i] = a[i] & b[i]
		for c[i] != 0 {
			c[i] &= c[i] -1
			n++
		}
	}
	return n
}