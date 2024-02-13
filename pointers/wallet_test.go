package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s , want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		// if got == nil {
		// 	t.Fatal("didn't get an error but wanted one")
		// }

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, 10)
		// got := wallet.Balance()
		// want := Bitcoin(10)

		// if got != want {
		// 	t.Errorf("got %s , want %s", got, want)
		// }
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, 10)
		assertError(t, err, nil)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ISF)
		assertBalance(t, wallet, startingBalance)

	})
}
