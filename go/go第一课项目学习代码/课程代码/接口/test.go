package main

import (
	"strings"
	"os"
	"io"
	"log"
	"fmt"
)

func main() {
	var i myint = 1
	test02(i)
}

type myint int

func (i myint)M1() {
	
}

func test01() {
    r := strings.NewReader("hello, gopher!\n")
    lr := io.LimitReader(r, 4)
    if _, err := io.Copy(os.Stdout, lr); err != nil {
        log.Fatal(err)
    }
}

func test02(i interface{}) {
	//i.M1()
	fmt.Println(i)
}