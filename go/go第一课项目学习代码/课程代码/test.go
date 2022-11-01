package main

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

func main() {
	test07()
}

type T struct{}

func (T) M1() {}
func (T) M2() {}

func (*T) M3() {}
func (*T) M4() {}

type I interface {
	M1()
	M2()
	M3()
}

type Str struct {
	I
}


func test01() {
	var i I
	var t T
	i = &t
	dumpMethodSet(&t)
	dumpMethodSet(i)
	var s Str
	
	dumpMethodSet(s)

}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

func test02() {

}

type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type t struct {
	a int
	b int
}

type S struct {
	*MyInt
	t
	io.Reader
	s string
	n int
}

func test03() {
	m := MyInt(17)
	r := strings.NewReader("hello, go")
	s := S{
		MyInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		Reader: r,
		s:      "demo",
	}

	var sl = make([]byte, len("hello, go"))
	s.Reader.Read(sl)
	fmt.Println(string(sl)) // hello, go
	s.MyInt.Add(5)
	fmt.Println(*(s.MyInt)) // 22
}

type E1 interface {
	M1()
	M2()
	M3()
}

type E2 interface {
	M1()
	M2()
	M4()
}

type T2 struct {
	E1
	E2
}

func (T2) M1() { println("T's M1") }
func (T2) M2() { println("T's M2") }

func test04() {
	t := T2{}
	t.E1.M1()
	t.E2.M2()
}

type T3 struct{}

func (T3) M1()  {}
func (*T3) M2() {}

type T4 T

func test05() {
	var t T3
	var pt *T3
	var t1 T4
	var pt1 *T4

	dumpMethodSet(t)
	dumpMethodSet(t1)

	dumpMethodSet(pt)
	dumpMethodSet(pt1)
}


type test2 struct {
	n int
	m int
}

type Test1 struct{}
func (Test1)Test() {}

type Itest1 interface {
	M1()
}

type Stest1 struct {
	Test1
	*test2
	Itest1
	a int
	b string
}

type Stest2 struct {
	T1 Test1
	t2 *test2
	I  Itest1
	a  int
	b  string
}

func test07() {
	var s1 Stest1
	var s2 Stest2
	dumpMethodSet(s1)
	dumpMethodSet(s2)
}