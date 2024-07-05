package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)
	numRounds := 3

	go throwingNinja(ch, numRounds)

	for i := 1; i <= numRounds; i++ {
		fmt.Println(<-ch)
	}
}

func throwingNinja(ch chan string, numRounds int) {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= numRounds; i++ {
		score := rand.Intn(10)

		ch <- fmt.Sprintf("Round %d: You scored %d", i, score)
	}
}
