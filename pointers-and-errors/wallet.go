package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is a sample custom type based on integer
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet has balance
type Wallet struct {
	balance Bitcoin
}

// Deposit adds amount to wallet's balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the balance of wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw reduces amount from wallet's balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("tried to withdraw more amount than available balance")
	}

	w.balance -= amount
	return nil
}