package main

import (
	"errors"
	"fmt"
	"time"

)

func main() {	
	test03()
    
}

func test01() {

	var ErrSentinel = errors.New("the underlying sentinel error")

	err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err1: %w", err1)
		println(err2 == ErrSentinel) //false
	if errors.Is(err2, ErrSentinel) {
		println("err2 is ErrSentinel")
		return
	}

	println("err2 is not ErrSentinel")
	
}

type MyError struct {
    e string
}
func (e *MyError) Error() string {
    return e.e
}
func test02() {
    var err = &MyError{"MyError error demo"}
    err1 := fmt.Errorf("wrap err: %w", err)
    err2 := fmt.Errorf("wrap err1: %w", err1)
    var e *MyError
    if errors.As(err2, &e) {
        println("MyError is on the chain of err2")
        println(e == err)                  
        return                             
    }                                      
    println("MyError is not on the chain of err2")

}

type field struct {
    name string
}

func (p *field) print() {
    fmt.Println(p.name)
}

func test03() {
    /*data1 := []*field{{"one"}, {"two"}, {"three"}}
    for _, v := range data1 {
        go v.print()
    }*/

    data2 := []field{{"four"}, {"five"}, {"six"}}
    for _, v := range data2 {
        go v.print()
    }

    time.Sleep(3 * time.Second)
}
