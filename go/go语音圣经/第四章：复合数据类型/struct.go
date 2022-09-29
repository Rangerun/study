package main

import (
	"fmt"
)
type E struct {
	id 	int
	name string
}

func getE(id int) *E {
	return &E{id:id, name:"angel"}
}

func main() {
	ptr1 := new(E)
	var ptr2 *E
	ptr1.id = 1
	ptr1.name ="angel"
	ptr2 = ptr1
	fmt.Println(ptr2.name)
}