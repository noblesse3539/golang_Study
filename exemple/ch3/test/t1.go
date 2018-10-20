package main

import (
	"math/cmplx"
	"fmt"
)

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Printf("%v \n", x)
	fmt.Println(x*y)
	fmt.Println(real(x*y))
	fmt.Println(imag(x*y))
	fmt.Println(1i*1i)
	
	fmt.Println(cmplx.Sqrt(-1))
	fmt.Println(cmplx.Abs(2+2i))
}