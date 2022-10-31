package main

import(
	"fmt"
	"unsafe"
	"encoding/json"
	"reflect"
)

func main()  {
	test05()
}

func test01() {	
	var sl1 []int
	var sl2 = []int{}
	fmt.Print("========基本区别=========\n")
	fmt.Printf("%v,len:%d,cap:%d,addr:%p\n", sl1, len(sl1), cap(sl1), &sl1)
	fmt.Printf("%v,len:%d,cap:%d,addr:%p\n", sl2, len(sl2), cap(sl2), &sl2)
	fmt.Printf("sl1==nil:%v\n", sl1 == nil)
	fmt.Printf("sl2==nil:%v\n", sl2 == nil)

	a1 := *(*[3]int)(unsafe.Pointer(&sl1))
	a2 := *(*[3]int)(unsafe.Pointer(&sl2))
	fmt.Print("========底层区别=========\n")
	fmt.Println(a1)
	fmt.Println(a2)

	type SliceDemo struct {
		Values []int
	}

	var s1 = SliceDemo{}
	var s2 = SliceDemo{[]int{}}
	bs1, _ := json.Marshal(s1)
	bs2, _ := json.Marshal(s2)
	fmt.Print("========序列化区别=========\n")
	fmt.Println(a1)
	fmt.Println(string(bs1))
	fmt.Println(string(bs2))
}

func test02() {
	var s = []int{1, 2}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s)) 
	fmt.Printf("0x%x\n", hdr.Data) // 0x10a30e0
	b := *(*[]int)(unsafe.Pointer(&hdr.Data))
	fmt.Println(b)
	fmt.Println(hdr.Cap)
}

func test03() {
	m := map[int]int{
        1: 11,
        2: 12,
        3: 13,
    }
	if _, ok := m[3]; ok {
		fmt.Println(11111)
	}
    fmt.Printf("{ ")
    for k, v := range m {
        fmt.Printf("[%d, %d] ", k, v)
    }
    fmt.Printf("}\n")
}

func test04() {
	
	type T struct {
		b byte
		i int64
		u uint16
	}
	
	type S struct {
		b byte
		u uint16
		i int64
	}
	
	var t T
	println(unsafe.Sizeof(t)) // 24
	var s S
	println(unsafe.Sizeof(s)) // 16

}

func test05() {
    var a = [5]int{1, 2, 3, 4, 5}
    var r [5]int

    fmt.Println("original a =", a)
	fmt.Printf("%p\n", &a[0])
    for i, v := range a {
		fmt.Printf("%p\n", &v)
        if i == 0 {
            a[1] = 12
            a[2] = 13
        }
        r[i] = v
    }
    fmt.Println("after for range loop, r =", r)
    fmt.Println("after for range loop, a =", a)
	a = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &a[0])
	for i, v := range a[:] {
		fmt.Printf("%p\n", &v)
        if i == 0 {
            a[1] = 12
            a[2] = 13
        }
        r[i] = v
    }

    fmt.Println("after for range loop, r =", r)
    fmt.Println("after for range loop, a =", a)
	
	/*var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			m["lucy"] = 24
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)*/
}

func test06() {
	var x interface{} = 13
    switch v := x.(type) {
    case nil:
        println("v is nil")
    case int:
        println("the type of v is int, v =", v)
    case string:
        println("the type of v is string, v =", v)
    case bool:
        println("the type of v is bool, v =", v)
    default:
        println("don't support the type")
    }
}