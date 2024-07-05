package main

import (
	"fmt"
	"time"
)

// In this example, we will learn to cases of using select statement
//
// case 1: use select statement without for loop
//   - In this case, which channel receives the message, that case will be called then stop the select statement
//
// case 2: use select statement with for loop
//   - In this case, all cases will be called because the select statement
//
// is inside the for loop.
// After a case inside the for loop is called, the for contine listening for other cases
func main() {
	ninja1, ninja2 := make(chan string), make(chan string)

	// Send values to channels in goroutines
	go captainElect(ninja1, "Ninja 1")
	go captainElect(ninja2, "Ninja 2")

	// in this scenario (case 1), "Ninja 1" or "Ninja 2" will be printed. not both
	// because when the main goroutine receives a message from which channel, it will print the message
	// then stop the select statement
	//
	// if use Sleep() inside captainElect(), the select will be blocked
	select {
	case msg1 := <-ninja1:
		fmt.Println(msg1)
	case msg2 := <-ninja2:
		fmt.Println(msg2)
	}

	// case 2
	roughlyFair()
}

func captainElect(ninja chan string, message string) {
	time.Sleep(time.Second * 3)
	ninja <- message
}

func roughlyFair() {

	// create 2 channels with no data and close them immediately
	ninja1 := make(chan interface{})
	close(ninja1)
	ninja2 := make(chan interface{})
	close(ninja2)

	var ninja1Count, ninja2Count int

	// because of using for, both channels will be listened
	// therefore, both ninja1Count and ninja2Count will be printed
	for i := 0; i < 1000; i++ {
		select {
		case <-ninja1: // this case is run when ninja1 closes
			ninja1Count++
		case <-ninja2:
			ninja2Count++
		}
	}

	fmt.Printf("Ninja1 count: %d, Ninja2 count: %d\n", ninja1Count, ninja2Count)
}
