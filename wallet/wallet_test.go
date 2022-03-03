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

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(got, want, t)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(got, want, t)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(wallet.Balance(), startingBalance, t)
		assertError(t, err, withdrawInsufficientFunds)
	})
}

func assertBalance(got Bitcoin, want Bitcoin, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("want an error but did't get one")
	}

	if err != want {
		t.Errorf("got %s, but want %s", err, want)
	}

}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("don't want an error but got one")
	}
}
