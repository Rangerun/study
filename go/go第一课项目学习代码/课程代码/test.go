package main

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

func main() {
	test06()
}

type T struct{}

func (T) M1() {}
func (T) M2() {}

func (*T) M3() {}
func (*T) M4() {}

type T1 struct{}

func (T1) M1() {}
func (T1) M2() {}

func test01() {
	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t T
	dumpMethodSet(t)
	dumpMethodSet(&t)
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
	type S = T1
	var t1 T1
	var s1 S
	dumpMethodSet(t1)
	dumpMethodSet(&t1)
	dumpMethodSet(s1)
	dumpMethodSet(&s1)
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

type T5 int
type t2 struct {
	n int
	m int
}

type I interface {
	M1()
}

type S1 struct {
	T1
	*t2
	I
	a int
	b string
}

type S2 struct {
	T1 T1
	t2 *t2
	I  I
	a  int
	b  string
}

func test06() {
    var s1 S1
    var s2 S2
    s1.t2.n = 1
    s1.M1()
    

}
