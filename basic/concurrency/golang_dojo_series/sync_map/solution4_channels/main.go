package main

import (
	"context"
	"sync"
	"time"
)

type CallCounter interface {
	Call(name string)
	Count(name string) int64
}

type CallCounterChannels struct {
	counters map[string]int64
	call     chan string
	count    chan string
	result   chan int64
}

func NewCallCounterChannels(ctx context.Context) *CallCounterChannels {
	c := CallCounterChannels{
		counters: make(map[string]int64),
		call:     make(chan string),
		count:    make(chan string),
		result:   make(chan int64),
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case name := <-c.call: // when the channel "call" receives a message, it will increment the counter int the map
				time.Sleep(1 * time.Second)
				c.counters[name]++
			case name := <-c.count: // when the channel "count" receives a message, it will send the counter value to the channel "result"
				c.result <- c.counters[name]
			}
		}
	}()

	return &c
}

func (c *CallCounterChannels) Call(name string) {
	// Send a message to the channel "call"
	c.call <- name
}

func (c *CallCounterChannels) Count(name string) int64 {
	// Send a message to the channel "count"
	c.count <- name

	// Receive a message from the channel "result"
	return <-c.result
}

func main() {
	counter := NewCallCounterChannels(context.Background())
	var wg sync.WaitGroup

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
