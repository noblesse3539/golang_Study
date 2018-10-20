package main

import (
	"strconv"
	"os"
	"fmt"
	"exercise/ch2/exer2_2/lenconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exer2_2: %v\n", err)
			os.Exit(1)
		}
		ft := lenconv.Feet(l)
		m := lenconv.Meter(l)
		
		fmt.Printf("%s = %s, %s = %s\n",
			ft, lenconv.FtToM(ft), m, lenconv.MToFt(m))
	}
}
