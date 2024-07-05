package main

import (
	"fmt"
	"sync"
)

//sync.RWMutex
// sync.RWMutex is a read-write mutual exclusion lock.
// It is a synchronization primitive that allows multiple goroutines to concurrently read shared data, but only one goroutine to write shared data
// at a time. This is useful for situations where you need to allow concurrent reads, but not concurrent writes to shared data.

// To use a sync.RWMutex, you must first lock it for reading and writing using the RLock and Lock methods, respectively.
// Once you are finished accessing the shared data, you must unlock the mutex using the RUnlock or Unlock methods, respectively.

// If another goroutine tries to lock the mutex for reading while it is locked for writing, it will be blocked until the mutex is unlocked.
// Similarly, if another goroutine tries to lock the mutex for writing while it is locked for reading, it will be blocked until the mutex is unlocked.

type CallCounter interface {
	Call(name string)
	Count(name string) int64
}

type CallCounterRWMutex struct {
	sync.RWMutex
	counters map[string]int64
}

func NewCallCounterRWMutex() *CallCounterRWMutex {
	c := CallCounterRWMutex{
		counters: make(map[string]int64),
	}

	return &c
}

func (c *CallCounterRWMutex) Call(name string, routineId int) {

	c.Lock()
	fmt.Printf("Goroutine %d is writing the map...\n", routineId)
	c.counters[name]++
	c.Unlock()
}

func (c *CallCounterRWMutex) Count(name string, routineId int) int64 {

	c.RLock()
	count := c.counters[name]
	c.RUnlock()

	return count
}

func main() {
	var wg sync.WaitGroup

	counter := NewCallCounterRWMutex()
	wg.Add(10)

	for i := 1; i <= 5; i++ {
		go func() {
			counter.Call("a", i)
			wg.Done()
		}()
	}

	for i := 6; i <= 10; i++ {
		go func() {
			val := counter.Count("a", i)
			fmt.Printf("Goroutine %d is reading the map... %d\n", i, val)
			wg.Done()
		}()
	}

	wg.Wait()

	println("DONE")

}
