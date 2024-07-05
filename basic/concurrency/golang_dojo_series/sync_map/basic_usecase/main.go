package main

import (
	"fmt"
	"sync"
)

// The overall purpose of this code is to demonstrate the concurrent usage of a sync.Map where multiple goroutines perform both write and read
// operations concurrently, showcasing how synchronization is handled by the sync.Map data structure.
func main() {
	// Create a new sync.Map
	var m sync.Map

	// Number of goroutines for concurrent operations
	numGoroutines := 5

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Writing goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			// Store key-value pair
			m.Store(id, id)

			// Load and print the value
			if value, ok := m.Load(id); ok {
				fmt.Printf("Goroutine %d: Key %d - Value %d\n", id, id, value)
			}
		}(i)
	}

	// Wait for all writing goroutines to finish
	wg.Wait()

	// Reading goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			// Load and print the value
			if value, ok := m.Load(id); ok {
				fmt.Printf("Reading Goroutine %d: Key %d - Value %d\n", id, id, value)
			} else {
				fmt.Printf("Reading Goroutine %d: Key %d not found\n", id, id)
			}
		}(i)
	}

	// Wait for all reading goroutines to finish
	wg.Wait()
}
