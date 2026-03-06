package main

import "fmt"
import "errors"

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var Error_Empty_Funds error = errors.New("there is not enough balance")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	current_balance := w.balance

	if current_balance - amount < 0 {
		return Error_Empty_Funds
	}

	w.balance -= amount

	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
