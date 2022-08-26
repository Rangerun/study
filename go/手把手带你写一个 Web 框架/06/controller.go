package main

import (
	"context"
	"coredemo/framework"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)
	durationCts, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1 * time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p:= recover(); p != nil {
				panicChan <- p
			}
		}()
		//time.Sleep(10 * time.Second)
		c.SetOkStatus().Json("ok")
		finish <- struct{}{}
	}()
	
	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.SetStatus(500).Json("panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCts.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.SetStatus(500).Json("time out")
		c.SetHasTimeout()
	}
	return nil

}