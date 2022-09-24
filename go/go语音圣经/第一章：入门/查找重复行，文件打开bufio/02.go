package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)		
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}


func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	count := 0
	for input.Scan() {
		if count > 4 {
			break;
		}
		count++
		counts[input.Text()]++		
	}
}