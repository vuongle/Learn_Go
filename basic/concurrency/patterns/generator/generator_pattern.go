package main

import (
	"fmt"
	"math/rand"
)

// The generator pattern is a way to create a function that gives us a series/stream of values on a channel.
// Think of it this way: we have a generator function that gives us something like a list we can pull values from, one at a time.
// When we get values from this list, we can use a for range in our code to handle each value as it comes.

// If running this example, the program will generate random numbers infinitely.
// because no value is sent to the done channel
func main() {
	done := make(chan int)
	defer close(done)

	// define a function that is passed to the generator function
	randNumFetcher := func() int {
		return rand.Intn(5000000000)
	}

	for rando := range generator(done, randNumFetcher) {
		fmt.Println(rando)
	}
}

// define generator function that receive a read-only channel, a function that return generic type T
// and returns a read-only channel that having generic type T
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
