package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"tcp/frame"
	"tcp/packet"

	"github.com/lucasepe/codename"
)

func main() {
	var wg sync.WaitGroup
	var num int = 5
	wg.Add(5)
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			startClient(i)
		}(i + 1)
	}
	wg.Wait()
}

func startClient(i int) {
	qiut := make(chan struct{})
	done := make(chan struct{})
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	
	defer conn.Close()
	fmt.Printf("[client %d]: dial ok", i)

	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}
	frameCode := frame.NewMyFrameCodec()
	vae counter int

	go func() {
		for {
			select {
			case <-quit:
				done <- struct{}{}
				return 
			default:
			}
			conn.SetReadDeadline(time.Now().Add(time.Second * 1))
			ackFramePayload, err := frameCodec.Decode(conn)
			if err != nil {
				if e, ok := err.(net.Error); ok {
					if e.Timeout() {
						continue
					}
				}
				panic(err)
			}
		}
	}

}
