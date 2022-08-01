package assessments

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

func (b bAccount) withdraw(n float64) bAccount {
	if b.balance >= n {
		b.balance -= n
	} else if b.balance < n || b.balance < 0 {
		fmt.Println("Unable to withdraw, balace is: ", b.balance)
	}
	return b
}
func (b bAccount) deposit(n float64) bAccount {
	b.balance += n
	return b
}
func (b bAccount) applyInterest() bAccount {
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
	return b
}
func (b bAccount) wire(c bAccount, n float64) (bAccount, bAccount) {
	if b.balance < n || b.balance < 0 {
		fmt.Println("Unable to withdraw, balace is: ", b.balance)
	} else {
		b = b.withdraw(n)
		c = c.deposit(n)
	}
	return b, c
}
func (e entity) changeAddress() entity {
	fmt.Print("Please enter new address: ")
	fmt.Scan(&e.address)
	return e
}
func (w wallet) displayAccounts() {
	for i := 0; i < len(w.accounts); i++ {
		if w.accounts[i].atype == "checking" {
			fmt.Print(w.accounts[i], " ")
		}
	}
	for i := 0; i < len(w.accounts); i++ {
		if w.accounts[i].atype == "investment" {
			fmt.Print(w.accounts[i], " ")
		}
	}
	for i := 0; i < len(w.accounts); i++ {
		if w.accounts[i].atype == "saving" {
			fmt.Print(w.accounts[i], " ")
		}
	}
	fmt.Println("")
}
func (w wallet) balance() {
	balance := 0.0
	for i := 0; i < len(w.accounts); i++ {
		fmt.Println(w.accounts[i])
	}
	for i := 0; i < len(w.accounts); i++ {
		balance += w.accounts[i].balance
	}
	fmt.Println("Overal Balance: ", balance)
}
func (w wallet) wire(b bAccount, c bAccount, n float64) (bAccount, bAccount) {
	if b.owner == w.owner {
		if b.balance < n || b.balance < 0 {
			fmt.Println("Unable to withdraw, balace is: ", b.balance)
		} else {
			b = b.withdraw(n)
			c = c.deposit(n)
		}
	}
	return b, c
}

func Mod6BankApp() {
	individualChecking := .01
	//individualInvestment := .02
	individualSaving := .05
	kat := entity{"kat", "534 Main Street", "Individual"}
	account1 := bAccount{"1235", kat, 500.00, individualChecking, "checking"}
	account2 := bAccount{"5648", kat, 500.00, individualSaving, "saving"}
	katW := wallet{}
	katW.id = "5264"
	katW.accounts = append(katW.accounts, account1)
	katW.accounts = append(katW.accounts, account2)
	katW.owner = kat
	fmt.Println(account1)
	fmt.Println(account2)
	account1, account2 = account1.wire(account2, 50)
	fmt.Println(account1)
	fmt.Println(account2)
}

//You will need to store the interest amount somewhere each time you apply interest to a particular account. Use appropriate data structures and logic to implement it.
