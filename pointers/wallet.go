package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

// type Stringer interface {
// 	String() string
// }

var ISF = errors.New("Insufficient funds!")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ISF
	}

	w.balance -= amount
	return nil
}
