package main

import (
	"errors"
	"fmt"
)

func main() {

}

func GenError() error {
	return errors.New("Demo error")

}

type Account struct {
	Balance float32
}

func (a *Account) Debit(amount float32) (float32, error) {
	if amount >= a.Balance {
		a.Balance = a.Balance - amount
		return a.Balance, nil
	}

	return 0, errors.New("Insufficient balance")
}

type BankingError struct {
	Code string
	Message string
}

func (b* BankingError)Error()string{
	fmt.Sprintf("Code:%s Message %s",b.code a.M)
}
