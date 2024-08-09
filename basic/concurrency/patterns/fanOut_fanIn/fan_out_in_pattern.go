package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

//The fan-out/fan-in pattern involves distributing tasks to multiple worker goroutines (fan-out) and then aggregating their results (fan-in).
// It’s useful for parallelizing tasks and combining their outcomes.
//
// This example use fan out and fan in pattern to resolve the problem in "generator_pipeline_combine.go".
// fan out: start multiple goroutines to read and process data from a channel
// fan in: combine data from multiple channels into a single channel
// See the pictures in the same folder as the flow

// If running this example, the program will generate random numbers in 10 times.
// After 10 times, it does not call repeatFunc() anymore
// Then the main function will exit, done channel will be closed -> 2 goroutines inside repeatFunc() and take() stop
func main() {
	done := make(chan int)
	defer close(done)

	// define a function that is passed to the generator function
	randNumFetcher := func() int {
		return rand.Intn(50000000)
	}
	randNumStream := generator(done, randNumFetcher) // stage 1

	// fan out
	CPUCount := runtime.NumCPU()
	// create a slice of un-buffered channels. each item in the slice is a un-buffered channel. CPUCount is the length of the slice
	primesFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primesFinderChannels[i] = primeFinder(done, randNumStream)
	}

	// fan in
	fanInStream := fanIn(done, primesFinderChannels...)
	for rando := range take(done, fanInStream, 10) {
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

func primeFinder(done <-chan int, randNumerStream <-chan int) <-chan int {
	isPrime := func(randNumber int) bool {
		for i := randNumber - 1; i > 1; i-- {
			if randNumber%i == 0 {
				return false
			}
		}

		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randNumber := <-randNumerStream:
				if isPrime(randNumber) {
					fmt.Printf("Prime Finder found: %d \n", randNumber)
					primes <- randNumber
				}
			}

		}
	}()

	return primes
}

func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fanInStream := make(chan T)

	// define an inner func to transer data from channels to fanInStream
	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fanInStream <- i:
			}
		}
	}

	// call transfer func for each channel
	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	// wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(fanInStream)
	}()

	return fanInStream
}
