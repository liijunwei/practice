package main

import (
	"errors"
	"fmt"
)

func main() {
	withdraw := makeWithdraw(100)
	var b int
	var err error
	b, err = withdraw(10)
	assert(err == nil)
	assert(b == 90)
	b, err = withdraw(80)
	assert(err == nil)
	assert(b == 10)
	b, err = withdraw(10)
	assert(err == nil)
	assert(b == 0)
	b, err = withdraw(1)
	assert(err != nil)
	assert(b == 0)
	fmt.Println("ok")
}

func makeWithdraw(balance int) func(amount int) (int, error) {
	return func(amount int) (int, error) {
		if amount > balance {
			return balance, errors.New("insufficient funds")
		}
		balance -= amount
		return balance, nil
	}
}

func assert(ok bool) {
	if !ok {
		panic("boom")
	}
}
