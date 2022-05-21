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
	capacity 		int
	tasks			chan Task
	active			chan struct{}
	quit			chan struct{}
	wg				sync.WaitGroup
}

func new(capacity int) *Pool {
	if (capacity <= 0) {
		capacity = 10
	}
	if (capacity > 100) {
		capacity = 50
	}

	p := &Pool{
		capacity: capacity,
		tasks: make(chan Task),
		active: make(chan struct{}, capacity),
		quit: make(chan struct{}),
	}
	p.run()
	return p
}

func (p *Pool) run() {
	idx := 0
	go func ()  {
		select {
		case <-p.quit:
			return 
		case p.active <- struct{}{}:
			idx++
			p.newWorker(idx)
		}
	}()
	return 
}

func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				println(err)
			}
			<-p.active
			p.wg.Done()
		}()
		
		for {
			select {
			case <-p.quit:
				<-p.active
				return 
			case t := <-p.tasks:
				println("todo")
				t()
			}
		}
	}()
}


func (p *Pool) Schedule(t Task) error {
	select {
	case <- p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) free() {
	close(p.quit)
	p.wg.Wait()
	return 
}













