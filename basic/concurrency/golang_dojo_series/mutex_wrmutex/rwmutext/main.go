package main

import (
	"fmt"
	"sync" // to import sync later on
)

type Account struct {
	lock    sync.RWMutex
	balance int
	Name    string
}

func (a *Account) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance -= amount
}

func (a *Account) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance += amount
}

func (a *Account) GetBalance() int {
	a.lock.RLock()
	defer a.lock.RUnlock()

	return a.balance
}

func main() {
	fmt.Println("Processing...")

	var account Account
	var wg sync.WaitGroup

	account.Name = "Test account"

	// deposit the account 20 times by 20 goroutines
	// the account is locked for each deposit -> the balance is 2000 after 20 deposits
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go account.Deposit(100, &wg)
	}

	// at this point, the balance is 2000
	// withdraw the account 10 times by 10 goroutines
	// the account is locked for each withdraw -> the balance is 1000 after 10 withdraw
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go account.Withdraw(100, &wg)
	}

	// at this point, the balance is 1000
	wg.Wait()
	fmt.Printf("Balance: %d\n", account.GetBalance())
}
