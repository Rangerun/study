package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)


func main() {
	test05()
}

func spawn(f func() error) <-chan error {
	c := make(chan error)

	go func() {
		c <- f()
	}()

	return c
}

func test01() {
	
	
	c := spawn(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("timeout")
	})
	fmt.Println(<-c)
	
}


func deadloop() {
    for {
    } 
}

func test02() {
	/*
	go1.13的话加上runtime.GOMAXPROCS(1) main goroutine在创建 deadloop goroutine 之后就无法继续得到调度
	但如果是go1.14之后的话即使加上runtime.GOMAXPROCS(1) main goroutine在创建 deadloop goroutine 之后还是可以得到调度，应该是因为增加了对非协作的抢占式调度的支持
	*/
	runtime.GOMAXPROCS(1)
    go deadloop()
    for {
        time.Sleep(time.Second * 1)
        fmt.Println("I got scheduled!")
    }

}


func produce(ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func test03() {
    ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()
	wg.Wait()
}

func worker() {
	println("worker is working")
	time.Sleep(time.Second)
}

func spawng(f func()) <-chan int {
	ch := make(chan int)
	go func () {
		f()
		ch <- 1
	}()
	return ch
}


func test04() {
	ch := spawng(worker)
	<-ch
	println("done")
}

type signal struct {}
func spawnGroup(f func(), n int, signalCh chan signal) <-chan signal{
	ch := make(chan signal)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			<-signalCh
			f()
			wg.Done()
		}()
	} 
	go func() {
		wg.Wait()
		ch <- signal{}
	}()
	return ch

}

func test05() {
	signalCh := make(chan signal)
	ch := spawnGroup(worker, 3, signalCh)
	time.Sleep(time.Second * 3)
	close(signalCh)
	<- ch
	println("done")
}