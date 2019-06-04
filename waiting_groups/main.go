package main

import (
	"fmt"
	"time"
	"sync"
)

func main () {
	// Add one waiting group for the following goroutine
	var waitingGroup sync.WaitGroup
	waitingGroup.Add(1)
	
	go func() {
		count(2)
		waitingGroup.Done()
	}()
	
	waitingGroup.Wait();
}

func count(by int) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i * by)
		time.Sleep(time.Millisecond * 500)
	}
}