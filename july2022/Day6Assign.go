package july2022

import "fmt"

type account struct {
	number  string
	balance float64
	name    string
	//transaction [][]int
}

func edit(a account) {

}
func search() {

}
func create() {

}

//sort by first name
//edits values in an account

func day6Asign() {
	a := new(account)
	a.number = "A5816"
	a.balance = 150
	a.name = "Bob Jack"

	fmt.Println(a)
	b := new(account)
	b.number = "A4484"
	b.balance = 200
	b.name = "John Doe"
	c := account{"A7156", 140, "Tim Shake"}
	fmt.Println(c)
}

//making a whole program apparently
// 1)Develop the Application with all functionality (Add ,alter ,search )
// 2)Time of each transection such as withdraw and deposit
// 3)Sorting based on balance or name
// 4)Trying to optimized your code
// 5)Evaluation based on algorithms and scalability
