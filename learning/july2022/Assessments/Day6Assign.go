package assessments

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type account struct {
	number  string
	balance float64
	owner   name
	//transaction [][]int
}
type name struct {
	first string
	last  string
}

func edit(a account) {

}
func search() {

}
func (a *account) create() {
	name1 := ""
	name2 := ""
	rand.Seed(time.Now().UnixNano())
	var random int = rand.Intn(1000) + 8999
	fmt.Println("account number: ", random)
	a.number = "A" + strconv.Itoa(random)
	fmt.Print("Please enter starting balance: ")
	fmt.Scan(&a.balance)
	fmt.Print("Please enter first and last name of owner: ")
	//spaces are counted as extra entries
	fmt.Scan(&name1, &name2)
	fmt.Println("first name: ", name1)
	fmt.Println("last name: ", name2)
	e := name{name1, name2}
	fmt.Println("Full Name: ", e)
	a.owner = e
}

//sort by first name
//edits values in an account

func Day6Asign() {
	// a := new(account)
	// a.number = "A5816"
	// a.balance = 150
	//a.owner = "Bob Jack"

	// fmt.Println(a)
	// b := new(account)
	// b.number = "A4484"
	// b.balance = 200
	//b.owner = "John Doe"
	// c := account{"A7156", 140}
	// fmt.Println(c)
	d := new(account)
	(*d).create()
	fmt.Println("New Account: ", d)

}

//making a whole program apparently
// 1)Develop the Application with all functionality (Add ,alter ,search )
// 2)Time of each transection such as withdraw and deposit
// 3)Sorting based on balance or name
// 4)Trying to optimized your code
// 5)Evaluation based on algorithms and scalability
