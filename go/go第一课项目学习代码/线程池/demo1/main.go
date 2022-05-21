package main

import (
	"fmt"
	"errors"
)

type MyError struct {
    error
}

var ErrBad = MyError{
    error: errors.New("bad things happened"),
}

func bad() bool {
    return false
}

func returnsError() error {
    var p *MyError = nil
    if bad() {
        p = &ErrBad
    }
    return p
}

func main() {
    err := returnsError()
	println(err)
    if err != nil {
        fmt.Printf("error occur: %+v\n", err)
        return
    }
    fmt.Println("ok")
}