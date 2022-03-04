package bank

import "errors"

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

/*
	Whitney's Attempts

	func Deposit(account *Account, amount float64) {
		account.Balance += amount
	}
*/
