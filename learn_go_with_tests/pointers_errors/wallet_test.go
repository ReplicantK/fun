package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Error("wanted:", want, "got:", got)		
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Error("expected:", want, "got:", got)
		}
	})

	t.Run("insufficient funds withdraw", func(t *testing.T) {
		starting_balance := Bitcoin(20)
		wallet := Wallet{starting_balance}

		get_balance := wallet.Balance()

		if starting_balance != get_balance {
			t.Error("expected the following balance:", starting_balance, "got instead:", get_balance)
		}

		err := wallet.Withdraw(Bitcoin(100))

		if err == nil {
			t.Fatal("expected an error, but did not receive one")
		}

		if err.Error() != Error_Empty_Funds.Error() {
			t.Error("expected: there is not enough balance, but got:", err.Error())
		}
	})
}
