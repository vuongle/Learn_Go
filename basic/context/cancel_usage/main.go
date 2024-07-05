package main

//
// This example shows how to use context to cancel a goroutine
//
import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	fmt.Println("Task Inprogress...")
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled")
	}
}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	go task(ctx)

	time.Sleep(time.Second * 1)
	cancel() // Send a signal to the ctx's Done channel

	time.Sleep(time.Second * 1)
}
