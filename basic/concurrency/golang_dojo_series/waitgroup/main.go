package main

import (
	"fmt"
	"sync"
)

func main() {
	var beeper sync.WaitGroup
	ninjas := []string{"Ninja 1", "Ninja 2", "Ninja 3"}

	beeper.Add(len(ninjas))
	for _, ninja := range ninjas {
		go attack(ninja, &beeper)
	}
	beeper.Wait()

	fmt.Println("main goroutine done")
}

func attack(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Killed the ninja at ", evilNinja)
	beeper.Done()
}
