package main

import (
	"context"
	"fmt"
	"time"
)

//
// This example shows how to use context to control deadline
//

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	go performTask(ctx)

	time.Sleep(3 * time.Second)
}

func performTask(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task completed or deadline exceeded:", ctx.Err())
		return
	}
}
