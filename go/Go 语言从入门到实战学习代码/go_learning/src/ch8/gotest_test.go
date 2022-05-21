package ch8

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)


func TestChannel(t *testing.T) {
	t.Log("befor:", runtime.NumGoroutine())
	t.Log(GetSingletonObj())
	t.Log("after:", runtime.NumGoroutine())
}