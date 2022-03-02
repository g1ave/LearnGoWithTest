package wallet

import (
	"errors"
	"fmt"
)

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
func (w *Wallet) Withdraw(mount Bitcoin) error {
	if mount > w.balance {
		return errors.New("oh no")
	}
	w.balance -= mount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
