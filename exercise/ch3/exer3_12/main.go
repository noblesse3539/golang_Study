package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagrams("asdffff","ffffdsa"))	// true
	fmt.Println(isAnagrams("asdff", "ffdss"))		// false
}

func isAnagrams(s1, s2 string) bool {
	n := len(s1)
	if n == len(s2) {
		for i := 0; i < n; i++ {
			if s1[i] != s2[n-1-i] {
				return false
			}
		}
		return true
	}
	return false
}
