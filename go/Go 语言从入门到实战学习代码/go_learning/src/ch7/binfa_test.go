package ch7

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string{
	time.Sleep(time.Microsecond * 10)
	return fmt.Sprintf("the first result %d", i)
}

func GetSingletonObj() string{
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestChannel(t *testing.T) {
	t.Log("befor:", runtime.NumGoroutine())
	t.Log(GetSingletonObj())
	t.Log("after:", runtime.NumGoroutine())
}




