package pointers

import "testing"

func TestWalleT(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalanceEqual(t, wallet.Balance(), Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalanceEqual(t, wallet.Balance(), Bitcoin(10))
	})

	t.Run("Withdraw larger than balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(250))
		assertBalanceEqual(t, wallet.Balance(), Bitcoin(100))
		assertError(t, err, ErrInSuffBalance)
	})
}

func assertBalanceEqual(t *testing.T, got Bitcoin, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("Err occurred")
	}
	if err != want {
		t.Errorf("got %s want %s", err.Error(), want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
