package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var i int32 // shared variable that is updated by multiple goroutines
	var wg sync.WaitGroup
	wg.Add(3)

	// go ProcessWithoutAtomic(&i, &wg)
	// go ProcessWithoutAtomic(&i, &wg)
	// go ProcessWithoutAtomic(&i, &wg)

	go ProcessWithAtomic(&i, &wg)
	go ProcessWithAtomic(&i, &wg)
	go ProcessWithAtomic(&i, &wg)

	wg.Wait()
	fmt.Println("i:", i)
}

// Without atomic, the shared variable is not locked. Therefore, multiple goroutines can inscrease the value at the same time.
// This result is the value is not the expected one.
// Theory:
// 3 go routines update the shared variable concurrently, each go routine inscreases the value by 2000 times.
// expected result: 3 * 2000 = 6000
// but the actual result is not.
func ProcessWithoutAtomic(variable *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2000; i++ {
		*variable++ // *: get value of the pointer
	}
}

func ProcessWithAtomic(variable *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2000; i++ {
		atomic.AddInt32(variable, 1) // the shared variable is locked here. no other go routine can update the value
	}
}
