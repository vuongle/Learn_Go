package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)

	go throwingNinja(ch)

	// use for...range to iterate over the channel
	// if not known how many elements are in the channel
	for msg := range ch {
		fmt.Println(msg)
	}
}

func throwingNinja(ch chan string) {
	rand.Seed(time.Now().UnixNano())
	numRounds := 3

	for i := 1; i <= numRounds; i++ {
		score := rand.Intn(10)

		ch <- fmt.Sprintf("Round %d: You scored %d", i, score)
	}

	// close the channel
	// if not use close(), it will cause deadlock in the main goroutine
	// to verify this, comment the close(ch) line and run the program
	close(ch)
}
