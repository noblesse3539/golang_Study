package main

import (
	"fmt"
)

func main() {
	c := [2]int{2, 1}
	a := [2]int{2, 1}

	d := c == a

	fmt.Println(d)
}
