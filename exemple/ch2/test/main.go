package main

import (
	"exemple/ch2/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F
}
