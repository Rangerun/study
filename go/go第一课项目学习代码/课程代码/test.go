package main

import (
	"fmt"
	"time"
	"errors"
)


func main() {
	test01()
}

func spawn(f func() error) <-chan error {
	c := make(chan error)

	go func() {
		c <- f()
	}()

	return c
}

func test01() {
	
	
	c := spawn(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("timeout")
	})
	fmt.Println(<-c)
	
}