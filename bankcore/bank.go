package bank

import (
	"errors"
	"fmt"
)

// Customer ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (account *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	account.Balance += amount
	return nil
}

func (account *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if account.Balance < amount {
		return errors.New("")
	}

	account.Balance -= amount
	return nil
}

func (account *Account) Statement() string {
	/*
		%v	the value in a default format
			when printing structs, the plus flag (%+v) adds field names
		%#v	a Go-syntax representation of the value
	*/
	return fmt.Sprintf("%v - %v - %v", account.Number, account.Name, account.Balance)
}

/*
	Whitney's Failed Attempts

	func Deposit(account *Account, amount float64) {
		account.Balance += amount
	}
*/
