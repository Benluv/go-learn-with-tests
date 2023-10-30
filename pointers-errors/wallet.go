package pointerserrors

import "fmt"

// in Go if a symbol (variables, types, functions et al) starts with a lowercase then it is private outside the package it is defined in
type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// no need to change Balance to use a pointer receiver. But it is convention for sake of consistency
func (w *Wallet) Balance() int {
	return w.balance
}
