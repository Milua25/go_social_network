package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, err error, want string) {
		t.Helper()
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if err.Error() != want {
			t.Errorf("got %q want %q", err.Error(), want)
		}
	}
	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient balance", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		withdrawAmount := Bitcoin(100)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(withdrawAmount)

		assertError(t, err, fmt.Sprintf("insufficient funds: have %v, want %v", startingBalance, withdrawAmount))
		assertBalance(t, wallet, startingBalance)
	})

	t.Run("new pointer", func(t *testing.T) {
		newPointer()
	})
}
