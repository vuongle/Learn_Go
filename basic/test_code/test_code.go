package main

import (
	"context"
	"fmt"
	"time"
)

// Main function
func main() {

	var i int8 = 120
	i += 10
	fmt.Println(i)
}

type User struct {
	name string
	age  int
}

func findUser(ctx context.Context, login string) (*User, error) {
	ch := make(chan *User)
	go func() {
		ch <- findUserInDB(login)
	}()

	select {
	case user := <-ch:
		return user, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout")
	}
}

func findUserInDB(login string) *User {
	time.Sleep(time.Second * 10)
	return &User{"John", 20}
}
