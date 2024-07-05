package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CallCounter interface {
	Call(name string)
	Count(name string) int64
}

type CallCounterMap struct {
	sync.Map
}

func NewCallCounterMap() *CallCounterMap {
	c := CallCounterMap{}

	return &c
}

func (c *CallCounterMap) Call(name string, id int) {
	value := int64(1)
	count, loaded := c.LoadOrStore(name, &value)
	if loaded {
		fmt.Printf("Goroutine #%d is writing to the map... \n", id)
		atomic.AddInt64(count.(*int64), int64(value))

	}
}

func (c *CallCounterMap) Count(name string) int64 {
	count, ok := c.Load(name)
	if ok {
		return atomic.LoadInt64(count.(*int64))
	}

	return -1
}

func main() {
	var wg sync.WaitGroup

	counter := NewCallCounterMap()
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func() {
			counter.Call("a", i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Count("a"))
}
