package wallet

import "testing"

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

	assert := func(got Bitcoin, want Bitcoin, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assert(got, want, t)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)
		assert(got, want, t)
	})
}
