package main

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
