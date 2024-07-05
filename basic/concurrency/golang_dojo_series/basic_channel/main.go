package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	signal := make(chan bool) // non-buffered channel
	ninja := "tommy"

	// start a goroutine and wait for it to finish
	go attack(ninja, signal)
	fmt.Println(<-signal) // Wait for a value is sent to the channel

	// Continue with the rest of the program once the goroutine has finished
	fmt.Println("main goroutine done")
}

func attack(target string, signal chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("Killed the ninja at ", target)

	// Send a signal to the channel to indicate that the goroutine has finished
	signal <- true
	signal <- false // this line does not make sense because after send "true" to the channel, the main channel contine running and execute the rest of the program
}
