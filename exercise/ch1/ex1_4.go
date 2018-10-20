package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	dupFiles := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupFiles, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupFiles, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, dupFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, dupFiles map[string]string, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		dupFiles[input.Text()] += arg
	}
	// NOTE: input.Err()에서의 잠재적 오류는 무시한다.
}
