package main

import (
	"fmt"
	"time"

	"workerpool2"
)

func main() {
	p := workerpool2.New(5, workerpool2.WithPreAllocWorkers(false), workerpool2.WithBlock(false))

	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
		})
		if err != nil {
			fmt.Printf("task[%d]: error: %s\n", i, err.Error())
		}
	}

	p.Free()
}
