package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

func main() {
	s := flag.String("s", "256", `set output format as one out of three; SHA256, SHA384, SHA512 
USAGE :
	-s 256 str...
	-s 384 str...
	-s 512 str...
	`)
	var str string
	flag.Parse()
	for _, i := range flag.Args() {
		str += i
	}
	if strings.Compare(*s, "256") == 0 {
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	} else if strings.Compare(*s, "384") == 0 {
		fmt.Printf("%x\n", sha512.Sum384([]byte(str)))
	} else if strings.Compare(*s, "512") == 0 {
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
	} else {
		fmt.Println("fatal!!")
	}
}
