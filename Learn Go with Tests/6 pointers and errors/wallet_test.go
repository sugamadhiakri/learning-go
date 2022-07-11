package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposti", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoErr(t, err)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but didn't get one.")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoErr(t testing.TB, got error) {
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
