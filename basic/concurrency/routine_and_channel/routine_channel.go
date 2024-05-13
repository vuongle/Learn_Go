package routine_and_channel

import (
	"fmt"
	"time"
)

func Run() {
	// create a channel having int data type by make() function
	ch := make(chan int)

	// start a new goroutine with anonymous function(no name)
	// this routine is used to add data to channel: add 10 data
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- i // assign a value to a channel
			fmt.Println("Add value to channel: ", i)
		}
		fmt.Println(len(ch))
	}()

	// receive data from channel (from main routine)
	// channel is first-in first-out
	// fmt.Println("value from channel: ", <-ch) // get value(first in) from channel -> print 1
	// fmt.Println("value from channel: ", <-ch) // get value(second in) from channel -> print 2
	// fmt.Println("value from channel: ", <-ch)

	fmt.Println("Progam ends")
}
