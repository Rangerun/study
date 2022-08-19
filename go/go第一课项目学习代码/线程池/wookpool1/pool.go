package workpool

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrNoIdleWorkerInPool = errors.New("no idle worker in pool") // workerpool中任务已满，没有空闲goroutine用于处理新任务
	ErrWorkerPoolFreed    = errors.New("workerpool freed")       // workerpool已终止运行
)

type Task func()

type Pool struct {
	active 	chan struct{}
	cap		int
	task    chan Task
	quit	chan struct{}
	wg      sync.WaitGroup
}

func New(cap int) *Pool {
	if cap <= 0 {
		return nil;
	}
	if cap > 1000 {
		return nil
	}
	pool := &Pool {
		active: make(chan struct{}, cap),
		cap: cap,
		task: make(chan Task),
		quit: make(chan struct{}),
	}

	pool.run()
	return pool
}

func (p Pool)newWork(i int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("发生错误");
			}
			p.wg.Done()
		}()
		for {
			select {
			case t := <- p.task:
				fmt.Printf("worker[%03d]: receive a task\n", i)
				t()
			case <- p.quit:
				return
			}
		}
	}()
}


func (p Pool)run() {
	var i int
	for {
		select {
			case p.active <- struct{}{}:
				i++
				p.newWork(i)
			case <-p.quit:
				return 
		}
	}
}


func (p Pool)Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed ;
	case p.task <- t:
		return nil
	}
}

func (p Pool)Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Printf("workerpool freed\n")
}