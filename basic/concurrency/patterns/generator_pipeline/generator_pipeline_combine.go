package main

import (
	"fmt"
	"math/rand"
)

// This example combines the generator pattern and the pipeline pattern.
// This continues from the "generator" example
// See the picture in the same folder as the flow

// If running this example, the program will generate random numbers in 10 times.
// After 10 times, it does not call repeatFunc() anymore
// Then the main function will exit, done channel will be closed -> 2 goroutines inside repeatFunc() and take() stop
func main() {
	done := make(chan int)
	defer close(done)

	// define a function that is passed to the generator function
	randNumFetcher := func() int {
		return rand.Intn(5000000000)
	}

	for rando := range take(done, generator(done, randNumFetcher), 10) {
		fmt.Println(rando)
	}
}

func generator[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			case stream <- fn(): // send a value return from fn to the stream channel
			}

		}
	}()

	return stream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)

	go func() {
		defer close(taken)

		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream: // this syntax means: read value from stream ("<-stream") then send it to taken channel ("taken <-")
			}
		}
	}()

	return taken
}
