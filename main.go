package main

import (
	"fmt"

	"github.com/ByungHaLee/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("me")
	account.Deposit(10)
	fmt.Println(account.Balance())
	account.Withdraw(20)
	fmt.Println(account.Balance())
}
