package main

import (
	"fmt"
	"time"
)

func main() {
	// create buffered channels for job queue and results
	jobsQueue := make(chan int, 100)
	results := make(chan int, 100)

	// create workers
	for i := 0; i < 3; i++ {
		go worker(i, jobsQueue, results)
	}

	// send 50 jobs to the jobs queue
	for i := 1; i <= 50; i++ {
		jobsQueue <- i
	}

	// close the jobs queue
	close(jobsQueue)

	// collect the results
	for i := 1; i <= 50; i++ {
		fmt.Printf("Result: %v\n", <-results)
	}

	// close the results channel to signal that all results received
	close(results)
}

// Define a worker:
//  1. Listen and receive(read) jobs from a queue. therefore, the format is "<-chan" (readonly channel)
//  2. Do some work
//  3. Send results to the results channel. therefore, the format is "chan<-" (write-only channel)
func worker(workerId int, jobsQueue <-chan int, results chan<- int) {
	for i := range jobsQueue {
		// #1
		fmt.Printf("worker %v received job: %v\n", workerId, i)

		time.Sleep(time.Second)

		// #2 and #3
		results <- i * i
	}
}
