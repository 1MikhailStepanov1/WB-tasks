package main

import (
	"fmt"
	"sync"
)

const goroutineNumber = 1000

type Increment struct {
	sync.Mutex
	value int
}

func main() {
	inc := new(Increment)
	wg := new(sync.WaitGroup)
	wg.Add(goroutineNumber)

	for i := 0; i < goroutineNumber; i++ {
		go func() {
			defer wg.Done()
			inc.Lock()
			inc.value++
			inc.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(inc.value)
}
