package main

import (
	"fmt"

	"github.com/ttaehyun/gonomad/banking/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
}
