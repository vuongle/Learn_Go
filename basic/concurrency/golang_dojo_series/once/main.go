package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionCompleted bool

func markMissionComplete() {
	fmt.Println("Marking mission as completed...")
	missionCompleted = true
}

func checkMissionComplete() {
	if missionCompleted {
		fmt.Println("Mission is now completed")
	} else {
		fmt.Println("Mission was failed")
	}
}

func foundTreasure() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}

// Logic:
//
// - Start 100 goroutines to find the treasure
// - Once the treasure is found, print "Mission is now completed"
// - When a goroutine found the treasure, other ones should be skipped
// - This means that finding the treasure should be done only one time
func main() {
	count := 100

	var wg sync.WaitGroup
	wg.Add(count)

	var once sync.Once

	for i := 0; i < count; i++ {
		go func() {
			if foundTreasure() {
				// The problem of the following line is that:
				// Even though a goroutine found the treasure, other ones still run, continue finding and marking as completed
				// Therefore, "Marking mission as completed..." is printed multiple times
				// This is not correct as the above logic.
				//markMissionComplete()

				// Solution:
				// Use once() to perform the markMissionComplete() only one time
				once.Do(markMissionComplete)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	checkMissionComplete()
}
