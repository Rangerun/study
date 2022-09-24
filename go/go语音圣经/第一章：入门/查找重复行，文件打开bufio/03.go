package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	counts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Printf("err is %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}