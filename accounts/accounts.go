package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw you don't have enough balance")

// NewAccount create account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
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
