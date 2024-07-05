package main

import "fmt"

func main() {
	// scenario 1
	// -------- not use regular map in concurrency mode (not in goroutines)---------
	// regularMap := make(map[int]interface{})
	// for i := 0; i < 10; i++ {
	// 	regularMap[0] = i
	// }
	// fmt.Println(regularMap)
	// =====> Run the above code -> no error occurred

	// scenario 2
	// -------- use regular map in concurrency mode (in goroutines)---------
	regularMap := make(map[int]interface{})
	for i := 0; i < 10; i++ {
		go func() {
			// all goroutines write same key
			regularMap[0] = i
		}()
	}
	fmt.Println(regularMap)
	// =====> Run the above code -> fatal error: concurrent map writes. Reason: regular map does not support concurrent writes

	// -------- Therefore, must use sync.Map in concurrency mode to resolve the above problem ---------
}
