package main

import (
	"fmt"
	"sync"
	"time"
)

// This example describes "Event Notification" usecase by using wait-broadcast

const maxWorkersCount = 10

func main() {
	var counter int32

	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	wg.Add(maxWorkersCount)

	// Start 10 go routines
	// The goroutines increment a counter and then either wait at the barrier or signal the barrier
	for i := range maxWorkersCount {
		go func(workerID int) {
			defer wg.Done()

			// Each worker goroutine performs its own work. In this example, it is Printf
			fmt.Printf("Worker %d performing work\n", workerID)
			time.Sleep(5 * time.Second) // Simulate work

			// then acquires the lock to increment the counter variable
			cond.L.Lock()
			defer cond.L.Unlock()
			counter++ // shared resource among goroutines

			// If the current worker is the last one to reach the barrier, it broadcasts the barrier condition using cond.Broadcast()
			// to wake up all waiting workers. Otherwise, it waits at the barrier using cond.Wait() to be notified by the last worker
			if counter == maxWorkersCount {
				fmt.Printf("Worker %d broadcast...\n", workerID)
				cond.Broadcast()
			} else {
				fmt.Printf("Worker %d is waiting at the barrier\n", workerID)
				cond.Wait()
			}

			fmt.Printf("Worker %d passed the barrier\n", workerID)
		}(i)
	}

	wg.Wait()
}
