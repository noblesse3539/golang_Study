package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "",""
	for temp,arg := range os.Args[0:] {
		s += sep + strconv.Itoa(temp) + arg
		sep = "\n"
	}
	fmt.Println(s)
}
