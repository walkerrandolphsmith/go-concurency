package main

import (
	"fmt"
	"time"
)

func main() {
	// synchronize goroutine without a explicit lock
	c := make(chan int)
	go count(2, c)

	// For each message received
	for msg := range c {
		fmt.Println(msg)
	}
}

func count (by int, c chan int) {
	for i := 1; i <=5; i++ {
		// send message through channel
		c <- by
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}