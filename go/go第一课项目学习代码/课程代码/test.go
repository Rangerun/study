package main

import (
	"fmt"
	"errors"
)



func main() {
	test06()
}

func test01() {	
	var i interface{} = 15 // ok
	fmt.Println(i)
	i = "hello, golang" // ok
	fmt.Println(i)
	type T struct{}
	var t T
	i = t  // ok
	i = &t // ok
}

func test02() {	
	var a int64 = 13
	var i interface{} = a
	v1, ok := i.(int64) 
	fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok) // v1=13, the type of v1 is int64, ok=true
	v2, ok := i.(string)
	fmt.Printf("v2=%s, the type of v2 is %T, ok=%t\n", v2, v2, ok) // v2=, the type of v2 is string, ok=false
	v3 := i.(int64) 
	fmt.Printf("v3=%d, the type of v3 is %T\n", v3, v3) // v3=13, the type of v3 is int64
	v4 := i.([]int) // panic: interface conversion: interface {} is int64, not []int
	fmt.Printf("the type of v4 is %T\n", v4) 
	
}


type MyInterface interface {
    M1()
}

type T int
               
func (T) M1() {
    println("T's M1")
}              
               
func test03() {  
    var t T    
    var i interface{} = t
    v1, ok := i.(MyInterface)
    if !ok {   
        panic("the value of i is not MyInterface")
    }          
    v1.M1()    
    fmt.Printf("the type of v1 is %T\n", v1) // the type of v1 is main.T
               
    i = int64(13)
    v2, ok := i.(MyInterface)
    fmt.Printf("the type of v2 is %T\n", v2) // the type of v2 is <nil>
    // v2 = 13 //  cannot use 1 (type int) as type MyInterface in assignment: int does not implement MyInterface (missing M1   method) 
}

func test04() {	
	var err error
	fmt.Printf("%T\n", err)
	err = errors.New("error1")
	fmt.Printf("%T\n", err)  // *errors.errorString
}


type QuackableAnimal interface {
    Quack()
}

type Duck struct{}

func (Duck) Quack() {
    println("duck quack!")
}

type Dog struct{}

func (Dog) Quack() {
    println("dog quack!")
}

type Bird struct{}

func (Bird) Quack() {
    println("bird quack!")
}                         
                          
func AnimalQuackInForest(a QuackableAnimal) {
    a.Quack()             
}                         
                          
func test05() {             
    animals := []QuackableAnimal{new(Duck), new(Dog), new(Bird)}
    for _, animal := range animals {
        AnimalQuackInForest(animal)
    }  
}



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
	fmt.Println(p)
    return p
}

func test06() {
	err := returnsError()
    if err != nil {
        fmt.Printf("error occur: %+v\n", err)
        return
    }
    fmt.Println("ok")
}