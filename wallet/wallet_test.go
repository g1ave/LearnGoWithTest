package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}

// add Withdraw method
func TestWithdraw(t *testing.T) {

	assertBalance := func(got Bitcoin, want Bitcoin, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error) {
		if err == nil {
			t.Error("want an error but did't get one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(got, want, t)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(got, want, t)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(wallet.Balance(), startingBalance, t)
		assertError(t, err)
	})
}
