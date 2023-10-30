package pointerserrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		testName := "deposit"
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, testName, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		testName := "withdraw"
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, testName, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		testName := "withdraw insufficient funds"
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, testName, wallet, startingBalance)
	})

}

func assertBalance(t testing.TB, name string, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	fmt.Printf("address of balance in test '%s' is %p \n", name, &wallet.balance)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
