package main

import (
	"fmt"
	"nomadcoder/banking"
)

func main() {
	account := banking.Account{Owner: "me", Balance: 1000}
	fmt.Println(account)
}
