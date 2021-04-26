package main

import (
	"fmt"

	"github.com/ByungHaLee/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("me")
	fmt.Println(*account)
}
