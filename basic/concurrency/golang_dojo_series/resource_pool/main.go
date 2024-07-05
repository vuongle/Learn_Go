package main

import (
	"fmt"
	"sync"
)

func main() {

	// create a pool that contains number of slices of 1024 bytes
	// each object in the pool is a slice of 1024 bytes
	var numMemPieces int
	memPool := &sync.Pool{
		New: func() interface{} {
			numMemPieces++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// create a bunch of works that access to the pool
	const numWorkers = 1024 * 1024 // = 1,048,576
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		// each goroutine will do:
		// Get an object from the pool, if not existed -> create a new one
		// then put the new object to the pool immediately
		// later goroutine will get the object from the pool and use it if
		// there is the object in the pool already
		// Therefore, number of objects in the pool are less then number of workers
		go func() {
			mem := memPool.Get().(*[]byte)

			fmt.Sprintln("taking some time to do on the resource")

			memPool.Put(mem)

			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("%d numMemPieces were created in the pool:\n", numMemPieces)
}
