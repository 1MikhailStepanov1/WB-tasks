package main

import (
	"fmt"
)

func main() {
	values := []interface{}{
		42,
		true,
		"privet",
		make(chan int),
		make(chan string),
		-666,
	}

	for _, v := range values {
		fmt.Println(detectType(v))
	}
}

func detectType(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan string, chan bool, chan interface{}:
		return "chan"
	default:
		return "unknown"
	}
}
