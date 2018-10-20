package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123456789.12345"))
	fmt.Println(comma("12345678.1234"))
	fmt.Println(comma("1234567.123"))
}

func comma(s string) string {
	n := len(s)
	for j := n - 1; j >= 0; j-- {
		if s[j] == '.' {
			n = j
			break
		}
	}
	i := 0
	var buf bytes.Buffer
	if n%3 > 0 {
		if n%3 > 1 {
			buf.WriteString(s[:2])
			i = 2
		} else {
			buf.WriteString(s[:1])
			i = 1
		}
		buf.WriteString(",")
	}
	for i < n {
		buf.WriteString(s[i : i+3])
		buf.WriteString(",")
		i += 3
	}
	result := buf.String()
	return result[:len(result)-1] + s[n:]
}
