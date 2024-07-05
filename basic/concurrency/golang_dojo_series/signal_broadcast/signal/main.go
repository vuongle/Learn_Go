package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Logic of this program:
// There are 2 routines: main and another
// The main waits for a signal to know it is ready to do something
// But the ready status is turned on by the another routine
// So, the main routine must use an infinite loop to check the ready variable

var ready bool

// Problem of the gettingReadyForMission() func (this func belongs to main routine):
// Must use a loop to check the ready variable that belongs to another routine
func gettingReadyForMission() {
	go gettingReady()
	workInternals := 0

	for !ready {
		workInternals++
	}
	fmt.Printf("We are now ready! After %d work internals", workInternals)
}

func gettingReady() {
	sleep()
	ready = true
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	sometime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(sometime)
}

// Solution:
// Use Cond
func gettingReadyForMissionWithCond() {
	cond := sync.NewCond(&sync.Mutex{}) // create a new condition variable in the main routine

	// Start a new routine to turn on the ready variable
	go gettingReadyWithCond(cond)

	workInternals := 0
	cond.L.Lock()
	for !ready {
		workInternals++
		cond.Wait() // after increasing the workInternals to 1, the main routine waits fot a signal from another routine. so, the loop is not continued
	}
	cond.L.Unlock()

	fmt.Printf("We are now ready! After %d work internals", workInternals)
}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal() // send a signal to the main routine that is waiting for the ready variable
}

func main() {
	//gettingReadyForMission()
	gettingReadyForMissionWithCond()
}
