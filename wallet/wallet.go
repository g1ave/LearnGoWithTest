package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

var withdrawInsufficientFunds = errors.New("cant withdraw, no enough funds")

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
func (w *Wallet) Withdraw(mount Bitcoin) error {
	if mount > w.balance {
		return withdrawInsufficientFunds
	}
	w.balance -= mount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
