package main

import (
	"context"
	"fmt"
	"time"
)

//
// This example shows how to use context to pass data between goroutines
//

func main() {

	// create a parent context
	ctx := context.Background()

	// use context.WithValue() to attach a user ID to the context
	ctx = context.WithValue(ctx, "UserID", 123)

	// passed the context to the performTask goroutine
	go performTask(ctx)

	time.Sleep(time.Second * 1)
}

func performTask(ctx context.Context) {

	// retrieves the user ID from the context
	userID := ctx.Value("UserID").(int) // type assertion
	fmt.Println("User ID:", userID)
}
