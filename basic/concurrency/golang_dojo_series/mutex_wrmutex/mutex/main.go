package main

import (
	"fmt"
	"sync" // to import sync later on
	"time"
)

// ------------------------ example of using mutex for a global variable ------------------------
var (
	lock   sync.Mutex
	rwLock sync.RWMutex
	GFG    = 0 // global variable accessed by multiple goroutines
)

// This is the function we’ll run in every goroutine.
func worker() {
	lock.Lock()

	GFG = GFG + 1

	lock.Unlock()
}

func basic() {
	// Launch several goroutines and increment
	for i := 0; i < 1000; i++ {
		go worker()
	}

	time.Sleep(time.Second)

	fmt.Println("Value of x", GFG)
}

// ------------------------ example of using mutex for a struct ------------------------
type Account struct {
	lock    sync.Mutex
	balance int
	Name    string
}

func (a *Account) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.lock.Lock()
	time.Sleep(time.Microsecond * 500)
	a.balance -= amount
	a.lock.Unlock()
}

func (a *Account) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.lock.Lock()
	time.Sleep(time.Microsecond * 500)
	a.balance += amount
	a.lock.Unlock()
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.balance
}

func account() {
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

func main() {
	// basic()
	account()
}
