package main

import (
	"fmt"
	"unicode/utf8"
	"reflect"
	"unsafe"
	"strings"
	"bytes"
)

func main() {
	//bar()
	strPtr()
	//testStr()
	//testJoin()
}

func bar() {    		
	var s = "中国人"
	fmt.Printf("the length of s = %d\n", len(s)) // 9
	fmt.Printf("0x%x\n", s[0]) // 0xe4：字符“中” utf-8编码的第一个字节
	for i := 0; i < len(s); i++ {
	fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
	}
	fmt.Printf("\n")

	var s1 = "中国人"
	fmt.Println("the character count in s is", utf8.RuneCountInString(s1)) 
	for _, c := range s {  
		fmt.Printf("0x%x ", c)
	}
}

func strPtr() {
    var s = "hello"
	fmt.Println(unsafe.Sizeof(s))
    hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // 将string类型变量地址显式转型为reflect.StringHeader
    fmt.Printf("0x%x\n", hdr.Data) // 0x10a30e0
    p := (*[5]byte)(unsafe.Pointer(hdr.Data)) // 获取Data字段所指向的数组的指针
	l := (int)(hdr.Len)
	fmt.Println(l)
	fmt.Println(*p)
}

func testStr() {
	var s string = "中国人"	
	// string -> []rune
	rs := []rune(s) 
	fmt.Printf("%x\n", rs) // [4e2d 56fd 4eba]			
	// string -> []byte
	bs := []byte(s) 
	fmt.Printf("%x\n", bs) // e4b8ade59bbde4baba
					
	// []rune -> string
	s1 := string(rs)
	fmt.Println(s1) // 中国人
					
	// []byte -> string
	s2 := string(bs)
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Printf("0x%x\n", hdr.Data)
	fmt.Printf("%p\n", &bs[0]) 
	fmt.Println(s2) // 中国人
}

func testJoin() {
	ss := []string{
        "A",
        "B",
        "C",
    }
 
    var b bytes.Buffer
    for _, s := range ss {
        fmt.Fprint(&b, s)
    }
 
    fmt.Println(b.String())

    var sb strings.Builder
    for _, s := range ss {
        fmt.Fprint(&sb, s)
    }
	var s1 = []byte{96, 97, 98}
	sb.Write(s1)
    fmt.Println(sb.String())
}

func unsafeEqual(a string, b []byte) bool {
    bbp := *(*string)(unsafe.Pointer(&b))
    return a == bbp
}