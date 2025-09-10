package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.ru.pool.ntp.org")
	if err != nil {
		log.Fatalf("Can't get current time: %s", err)
	}
	fmt.Println(time) //nolint:forbidigo
}
