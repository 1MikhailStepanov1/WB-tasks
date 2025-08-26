package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeToStop := flag.Int("time", 3, "seconds before stop")
	flag.Parse()

	timer := time.NewTimer(time.Duration(*timeToStop) * time.Second)

	connection := make(chan int)

	go func() {
		for {
			select {
			case <-timer.C:
				close(connection)
				return
			default:
				connection <- rand.Intn(100)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	for data := range connection {
		fmt.Println(data)
	}
}
