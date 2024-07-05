package main

import "sync"

// sync.Mutex
// sync.Mutex is a mutual exclusion lock. It is a synchronization primitive that allows concurrent goroutines to access shared data safely.
// A mutex ensures that only one goroutine can access the shared data at a time, preventing data races and other concurrency problems.
// To use a sync.Mutex, you must first lock it before accessing the map to shared data.
// Once you are finished accessing the shared data, you must unlock the mutex.
// If another goroutine tries to lock the mutex while it is already locked, it will be blocked until the mutex is unlocked.

type CallCounter interface {
	Call(name string)
	Count(name string) int64
}

type CallCounterMutex struct {
	sync.Mutex // this way embeds the mutex struct in this struct (a kind of composition)
	counters   map[string]int64
}

func NewCallCounterMutex() *CallCounterMutex {
	c := CallCounterMutex{
		counters: make(map[string]int64),
	}

	return &c
}

func (c *CallCounterMutex) Call(name string) {
	c.Lock()
	c.counters[name]++
	c.Unlock()
}

func (c *CallCounterMutex) Count(name string) int64 {
	c.Lock()
	count := c.counters[name]
	c.Unlock()

	return count
}

func main() {

	var wg sync.WaitGroup
	counter := NewCallCounterMutex()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			counter.Call("a")
			wg.Done()
		}()
	}

	wg.Wait()

	println(counter.Count("a"))
}
