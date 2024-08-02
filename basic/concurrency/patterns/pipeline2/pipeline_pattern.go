package main

import (
	"fmt"
	"time"
)

// This is another example of pipeline. the logic of pipeline is that:
// stage 1: convert slice to channel
// stage 2: square the number
// stage 3: display the number in string
func main() {
	// input
	numbers := []int{2, 3, 4, 7, 1}

	// stage 1
	dataChannel := sliceToChannel(numbers)

	// stage 2
	// output of stage 1 is input of stage 2
	finalChannel := sq(dataChannel)

	// stage 3
	// output of stage 2 is input of stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}

// This func returns a read-only channel, therefore the format is "<-chan"
// This func returns the "out" channel while the inside goroutine is still running
func sliceToChannel(nums []int) <-chan int {

	// create an unbuffered channel
	out := make(chan int)

	go func() {
		for _, n := range nums {
			// right after sending a value to the "out" channel -> this goroutine blocks
			// and waits for another goroutine (the goroutine in sq() func) to read the value
			time.Sleep(time.Second)
			out <- n
		}

		close(out)
	}()

	return out
}

// This func receives and returns read-only channels, therefore the format is "<-chan"
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			// right after reading a value to the "in" channel -> this goroutine blocks
			// and waits for another goroutine (the goroutine in sliceToChannel() func) to send the value
			out <- n * n
		} // for stops when the "in" channel is closed (the channel is closed in sliceToChannel() func)

		close(out)
	}()

	return out
}
