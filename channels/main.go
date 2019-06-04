package main

import (
	"fmt"
	"time"
)

func main () {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			// send message through channel
			c1 <- "500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			// send message through channel
			c2 <- "2s"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		// Allow for the first receievd message to be processed
		// i.e. c1 to not wait on the slower channel c2
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
	}

}