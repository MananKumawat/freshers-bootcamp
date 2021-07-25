package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type BankAccount struct { // Struct for Bank Account
	accountnumber string
	balance int
	mutex sync.Mutex
}

func (bankaccount *BankAccount) AddMoney(amount int, wg *sync.WaitGroup) { // Function for adding money in bank account
	defer wg.Done()
	bankaccount.mutex.Lock()
	fmt.Println("Adding", amount, "Rs. to account")
	if amount < 0 {
		err := errors.New("Entered Negative Amount")
		log.Fatal(err)
	}

	bankaccount.balance += amount
	fmt.Println("Current Balance:", bankaccount.getbalance())
	bankaccount.mutex.Unlock()
}

func (bankaccount *BankAccount) SubtractMoney(amount int, wg *sync.WaitGroup) { // Function for removing money from bank account
	defer wg.Done()
	bankaccount.mutex.Lock()
	fmt.Println("Removing", amount, "Rs. from account")
	if amount < 0 {
		err := errors.New("Entered Negative Amount")
		log.Fatal(err)
	}
	if amount > bankaccount.getbalance() {
		err := errors.New("Insufficient Balance")
		log.Fatal(err)
	}

	bankaccount.balance -= amount
	fmt.Println("Current Balance:", bankaccount.getbalance())
	bankaccount.mutex.Unlock()
}

func (bankaccount BankAccount) getbalance() int {
	return bankaccount.balance;
}

func main() {
	rand.Seed(time.Now().UnixNano())
	bankaccount := BankAccount{"2030xxxx", 500, sync.Mutex{}} // Initializing a bank account with 500 Rs.

	fmt.Println("Initial Balance:", bankaccount.getbalance())

	var wg sync.WaitGroup

	for cases:=0; cases<3; cases++ {
		amount := rand.Intn(20) * 100
		if rand.Intn(2) == 0 {
			wg.Add(1)
			go bankaccount.AddMoney(amount, &wg)
		} else {
			wg.Add(1)
			go bankaccount.SubtractMoney(amount, &wg)
		}
	}

	wg.Wait()
}