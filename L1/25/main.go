package main

import (
	"fmt"
	"time"
)

func sleep(dur time.Duration) {
	<-time.After(dur)
}

func main() {
	fmt.Println(time.Now())
	sleep(3 * time.Second)
	fmt.Println(time.Now())
}
