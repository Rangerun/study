package main

import (
    "fmt"
    "time"
)

type field struct {
    name string
}

func (p *field) print() {
    fmt.Println(p.name)
}

func main() {
    /*data1 := []*field{{"one"}, {"two"}, {"three"}}
    for _, v := range data1 {
        go v.print()
    }*/

    data2 := []field{{"four"}, {"five"}, {"six"}}
    for _, v := range data2 {
		p := &v
		fmt.Println("---%s", &p);
        //go v.print()
		//go (*field).print(&v)
    }

    time.Sleep(3 * time.Second)
}