package main

import (
	_ "test01/internal/store"
	"test01/store/factory"
	
)

func main() {
	s, err := factory.New("mem")
	
	if err != nil {
		panic(err)
	}
	if s == nil {
		panic(err)
	}

}
