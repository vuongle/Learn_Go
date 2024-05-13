package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHi(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Use WaitGroup to wait all routines finish before main finishes
var wg sync.WaitGroup

func go1() {
	for i := 1; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Logic %d in go1\n", i)
	}
	wg.Done() // indicate that routine "go1" is done. If there is no this line, a deadlock will happen
}

func go2() {
	for i := 1; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Logic %d in go2\n", i)
	}
	wg.Done() // indicate that routine "go2" is done
}

func main() {

	// -----------------------------------
	// Simple goroutine
	// -----------------------------------

	// starts a new goroutine running "sayHi()" function with the keyword "go"
	//go sayHi("AAA")
	//sayHi("BBB") // this belongs to main routine

	// -----------------------------------
	// Synchonized goroutine
	// -----------------------------------
	fmt.Println("Logic in main")

	wg.Add(2)
	go go1()
	go go2()
	wg.Wait()
	fmt.Println("Progam ends")

	// -----------------------------------
	// routine and channels
	// -----------------------------------

}

//https://www.youtube.com/watch?v=y8_YRzXFQ84&list=PLVDJsRQrTUz5icsxSfKdymhghOtLNFn-k&index=4
