package main

type E struct {
	id 	int
	name string
}

func getE(id int) *E {
	return &E{id:id, name:"angel"}
}

func mian() {
	id := 1
	getE(id).name = "0"
}