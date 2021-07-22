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

func (bankaccount *BankAccount) AddMoney(amount int) { // Function for adding money in bank account
	fmt.Println("Adding", amount, "Rs. to account")
	if amount < 0 {
		err := errors.New("Entered Negative Amount")
		log.Fatal(err)
	}

	bankaccount.mutex.Lock()
	bankaccount.balance += amount
	bankaccount.mutex.Unlock()
	fmt.Println("Current Balance:", bankaccount.balance)
}

func (bankaccount *BankAccount) SubtractMoney(amount int) { // Function for removing money from bank account
	fmt.Println("Removing", amount, "Rs. from account")
	if amount < 0 {
		err := errors.New("Entered Negative Amount")
		log.Fatal(err)
	}
	if amount > bankaccount.balance {
		err := errors.New("Insufficient Balance")
		log.Fatal(err)
	}

	bankaccount.mutex.Lock()
	bankaccount.balance -= amount
	bankaccount.mutex.Unlock()
	fmt.Println("Current Balance:", bankaccount.balance)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	bankaccount := BankAccount{"2030xxxx", 500, sync.Mutex{}} // Initializing a bank account with 500 Rs.

	fmt.Println("Initial Banalce:", bankaccount.balance)

	for cases:=0; cases<3; cases++ {
		amount := rand.Intn(20) * 100
		if rand.Intn(2) == 0 {
			bankaccount.AddMoney(amount)
		} else {
			bankaccount.SubtractMoney(amount)
		}
	}
}