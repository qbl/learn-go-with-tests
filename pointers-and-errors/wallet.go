package wallet

// Bitcoin is a sample custom type based on integer
type Bitcoin int

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
