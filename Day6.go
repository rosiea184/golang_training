package main

import "fmt"

type account struct {
	number  string
	balance float64
	name    string
	typeA   string
}

func main() {
	a := new(account)
	a.number = "A5816"
	a.balance = 150
	a.name = "Bob"
	a.typeA = "checking"

	fmt.Println(a)

}
