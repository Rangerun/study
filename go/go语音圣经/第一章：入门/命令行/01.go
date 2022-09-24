package main

import (
	"os"
	"fmt"
)

func main() {
	s := ""
	for i := 0; i < len(os.Args); i++ {
		s += "_" + os.Args[i]
	}
	fmt.Println(s)
	return 
}