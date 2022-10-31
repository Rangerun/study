package main

import (
    "bytes"
    "fmt"
    "strings"
    "testing"
)

/*func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}*/

/*
func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(1000000)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(1000000)
	}
}*/

 
func BenchmarkBuffer(b *testing.B) {
    var buf bytes.Buffer
    for i := 0; i < b.N; i++ {
        fmt.Fprint(&buf, "?")
        _ = buf.String()
    }
}
 
func BenchmarkBuilder(b *testing.B) {
    var builder strings.Builder
    for i := 0; i < b.N; i++ {
        fmt.Fprint(&builder, "?")
        _ = builder.String()
    }
}
