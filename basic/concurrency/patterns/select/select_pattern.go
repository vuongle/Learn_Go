package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// select pattern:
// allow to listen on multiple channels. It will wait until one of the channels has a value.
// which channel has a value first -> corresponding case will be executed
func main() {
	//simpleSelectPattern()

	//forSelect()

	forSelectDonePattern()
}

func simpleSelectPattern() {

	// create 2 unbuffered channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// start 2 goroutines to send data to 2 channels
	go func() {
		rand := rand.IntN(5)
		time.Sleep(time.Duration(rand) * time.Second)
		ch1 <- "data in channel 1"
	}()

	go func() {
		rand := rand.IntN(5)
		time.Sleep(time.Duration(rand) * time.Second)
		ch2 <- "data in channel 2"
	}()

	// use select to listen on multiple channels
	// select will block the program until one of the channels has a value
	// which channel has a value first -> that case will be executed
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	fmt.Println("done")
}

func forSelect() {
	chars := []string{"a", "b", "c"}

	// create buffered channel
	ch := make(chan string, 3)

	// one way to send values to channel without using select
	// for _, s := range chars {
	// 	ch <- s
	// }

	// using select to send values to channel
	for _, s := range chars {
		select {
		case ch <- s:
		}
	}

	// close the channel so that the main goroutine can read from it
	// without closing -> cause deadlock
	close(ch)

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("done")
}

func forSelectDonePattern() {
	done := make(chan bool)

	go doWork(done)

	// send done value to the "done" channel after 5s by closing the channel or by sending a value
	time.Sleep(5 * time.Second)
	//close(done)
	done <- true

	fmt.Println("done")
}

// done: this is a read-only channel, therefore the format is "<-chan".
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Println("doing work")
		}
	}
}
