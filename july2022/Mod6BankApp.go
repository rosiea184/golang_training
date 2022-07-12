package july2022

import "fmt"

type bAccount struct {
	number   string
	owner    entity
	balance  float64
	interest float64
	atype    string
}
type entity struct {
	id      string
	address string
	atype   string
}
type wallet struct {
	id       string
	accounts []bAccount
	owner    entity
}

func (b bAccount) withdraw(n float64) {
	if b.balance >= n {
		b.balance -= n
	} else if b.balance < n || b.balance < 0 {
		fmt.Println("Unable to withdraw, balace is: ", b.balance)
	}
}
func (b bAccount) deposit(n float64) {
	b.balance += n
}
func (b bAccount) applyInterest() {
	if b.owner.atype == "individual" {
		if b.atype == "checking" {
			b.balance += (b.balance * .01)
		} else if b.atype == "saving" {
			b.balance += (b.balance * .05)
		} else if b.atype == "investment" {
			b.balance += (b.balance * .02)
		}

	} else if b.owner.atype == "buisness" {
		if b.atype == "checking" {
			b.balance += (b.balance * .005)
		} else if b.atype == "saving" {
			b.balance += (b.balance * .02)
		} else if b.atype == "investment" {
			b.balance += (b.balance * .01)
		}
	}
}
func (b bAccount) wire(c bAccount, n float64) {
	b.withdraw(n)
	c.deposit(n)
}
func (e entity) changeAddress() {
	fmt.Print("Please enter new address: ")
	fmt.Scan(&e.address)
}
func (w wallet) displayAccounts() {
	for i:=0; i < len(w.accounts); i++(
		if w.accounts[i].atype == "checking"{
			fmt.Print(w.accounts[i], " ")
		}
	)
	for i:=0; i < len(w.accounts); i++(
		if w.accounts[i].atype == "investment"{
			fmt.Print(w.accounts[i], " ")
		}
	)
	for i:=0; i < len(w.accounts); i++(
		if w.accounts[i].atype == "saving"{
			fmt.Print(w.accounts[i], " ")
		}
	)
	fmt.Println("")
}
func (w wallet) balance()         {
	balance := 0.0
	for i:=0; i < len(w.accounts); i++{
		fmt.Println(w.accounts[i])
	}
	for i:=0; i < len(w.accounts); i++{
		balance += w.accounts[i].balance
	} 
	fmt.Println("Overal Balance: ", balance)
}
func (w wallet) wire()            {}

func Mod6BankApp() {

}

//You will need to store the interest amount somewhere each time you apply interest to a particular account. Use appropriate data structures and logic to implement it.
