package wallet

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(mount Bitcoin) {
	w.balance += mount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw method of Wallet
func (w *Wallet) Withdraw(mount Bitcoin) {
	w.balance -= mount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
