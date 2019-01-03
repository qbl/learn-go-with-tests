package wallet

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		actual := wallet.Balance()

		if actual != expected {
			t.Errorf("actual: %s; expected: %s;", actual, expected)
		}
	}

	assertError := func(t *testing.T, actual error, expected string) {
		t.Helper()
		if actual == nil {
			t.Fatal("should returns error but doesn't")
		}

		if actual.Error() != expected {
			t.Errorf("actual: '%s'; expected: '%s'", actual, expected)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with insufficient amount", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "tried to withdraw more amount than available balance")
	})
}