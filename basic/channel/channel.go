package main

import (
	"fmt"
	"time"
)

// There are 2 kinds of channel:
//
//	Unbuffered channel: have no size
//	Bufferred channel: Have size
func main() {
	// --------- Unbuffered channel
	ch := make(chan int)

	// start a new goroutine
	go func() {
		time.Sleep(time.Second * 2)
		// send a value to channel inside a goroutine
		ch <- 100 // at this point, the program is blocked, untill another goroutine get "100" out of the channel
		fmt.Println("SENT")
	}()

	fmt.Println("START READING")
	// read a value from a channel inside another goroutine
	fmt.Println(<-ch)
	fmt.Println("DONE")

	// --------- buffered channel
	bCh := make(chan int, 2) //2: size of the channel
	bCh <- 1                 // at this point, the program is NOT blocked
	bCh <- 2                 // at this point, the program is NOT blocked either

	fmt.Println(<-bCh)
	fmt.Println(<-bCh)
}
