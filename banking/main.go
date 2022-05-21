package main

import (
	"fmt"

	"github.com/ttaehyun/gonomad/banking/banking_module"
)

func main() {
	account := banking.Account{Owner: "me", Balance: 100}
	fmt.Println(account)
}
