package wallet

// Wallet has balance
type Wallet struct {
	balance int
}

// Deposit adds amount to wallet's balance
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance returns the balance of wallet
func (w *Wallet) Balance() int {
	return w.balance
}
