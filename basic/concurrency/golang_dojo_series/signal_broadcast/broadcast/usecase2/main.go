package main

import (
	"fmt"
	"sync"
	"time"
)

// This example is same as the "usecase1", but it is more easier to understand

func dummyGoroutine(routineId int, cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()
	fmt.Printf("Goroutine %d is waiting...\n", routineId)

	// wait for signal and release the lock so that other goroutines can acquire the lock and execute
	cond.Wait()
	fmt.Printf("Goroutine %d received the signal.\n", routineId)
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go dummyGoroutine(1, cond)
	go dummyGoroutine(2, cond)

	time.Sleep(1 * time.Second)

	// main routine broadcasts to all goroutines
	cond.Broadcast()
	time.Sleep(1 * time.Second)
}
