package main

import (
	"fmt"
	"time"
)

func sayHello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go sayHello("world")
	sayHello("hello")
}
