package pointerserrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

// in Go if a symbol (variables, types, functions et al) starts with a lowercase then it is private outside the package it is defined in
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// if pointer removed from wallet, then the memory location of the wallet changes
	fmt.Printf("address of balance in Deposit is %p when not pointed at value\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.balance -= amount
	return nil
}

// no need to change Balance to use a pointer receiver. But it is convention for sake of consistency
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
