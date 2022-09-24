package main

import "fmt"

func main() {
	var i *int
	i = test()
	fmt.Println(*i)
	a := newI()
	*a = 2
	fmt.Println(*a)
}

func test() *int {
	v := 1
	return &v 
}

func newI() *int {
	return new(int)
}