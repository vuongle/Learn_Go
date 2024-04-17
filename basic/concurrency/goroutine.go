package main

import (
	"fmt"
	"time"
)

func sayHi(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	// starts a new goroutine running "sayHi()" function with the keyword "go"
	go sayHi("AAA")
	sayHi("BBB") // this belongs to main routine

	// ------------------------- routine and channels---------------------------------------
	// create a channel by make() function
	ch := make(chan int)

	// start a new goroutine with anonymous function(no name)
	// this routine is used to add data to channel
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Second * 1)
			ch <- i // assign a value to a channel
		}
	}()

	// receive data from channel (from main routine)
	// channel is first-in first-out
	fmt.Println("channel: ", <-ch) // get value from channel once -> print once
	fmt.Println("channel: ", <-ch)
}
