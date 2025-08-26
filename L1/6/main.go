package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	//Finish by condition
	wg.Add(1)
	stopFlag := false
	go func() {
		defer wg.Done()
		for !stopFlag {
			fmt.Println("Goroutine 1 is running...")
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Goroutine 1 stopped")
	}()
	stopFlag = true

	//Finish by channel signal
	wg.Add(1)
	notify := make(chan struct{})
	go func() {
		defer wg.Done()
		<-notify
		fmt.Println("Goroutine 2 stopped")
	}()
	notify <- struct{}{}

	//Finish by context
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		defer wg.Done()
		<-ctx.Done()
		fmt.Println("Goroutine 3 stopped")
	}(ctx)
	cancel()

	//Finish by runtime.Goexit
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Goroutine 4 stopped")
		runtime.Goexit()
	}()

	wg.Wait()
}
