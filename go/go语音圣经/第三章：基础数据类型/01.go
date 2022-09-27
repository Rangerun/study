package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(basename1("a/b/c.go")) // "c"
	//fmt.Println(basename1("c.d.go"))   // "c.d"
	//fmt.Println(basename1("abc"))      // "abc"
	fmt.Println(comma1("1234567"))      // "abc"
	fmt.Println(intsToString([]int{1, 2, 3}))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i + 1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename1(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma1(s string) string {
	var buf bytes.Buffer
	
	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		
		if i % 3 == 0 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}


func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
