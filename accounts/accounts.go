package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw you don't have enough balance")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (account *Account) Deposit(amount int) {
	account.balance += amount
}

func (account *Account) Balance() int {
	return account.balance
}

func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		return errNoMoney
	}
	account.balance -= amount
	return nil
}

func (account *Account) ChangeOwner(newOwner string) {
	account.owner = newOwner
}

func (account Account) Owner() string {
	return account.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
