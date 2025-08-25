package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomString(length int) string {
	res := make([]byte, length)
	for i := range res {
		res[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(res)
}

func work(id int, input <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range input {
		fmt.Printf("Worker %d, data: %s\n", id, data)
	}

}

func main() {
	workersAmount := flag.Int("workers", 10, "number of workers")
	flag.Parse()

	input := make(chan string)
	wg := new(sync.WaitGroup)

	for i := 0; i < *workersAmount; i++ {
		wg.Add(1)
		go work(i, input, wg)
	}

	for {
		input <- generateRandomString(rand.Intn(15))
		time.Sleep(time.Millisecond * 100)
	}
}
