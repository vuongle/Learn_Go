package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

//
// This example shows how to use context to control timeout
// Launch multiple goroutines to fetch data from different APIs concurrently.
// If any of the API requests exceed the timeout duration, the context's cancellation signal is propagated, canceling all other ongoing requests.
//

func main() {

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://api.example.com/users",
		"https://api.example.com/products",
		"https://api.example.com/orders",
	}

	// Create a channel to store the results
	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}

// syntax "chan<-"" means that the channel can only be written to
func fetchAPI(ctx context.Context, url string, results chan<- string) {

	// Create a new request that will be canceled after 5s because the context is canceled after 5s
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error fetching %s: %s", url, err.Error())
		return
	}

	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}
