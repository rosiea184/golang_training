package july2022

import "fmt"

type testaccount struct {
	number  string
	balance float64
	name    string
	typeA   string
}

func Day6() {
	//Create a struct
	a := new(testaccount)
	a.number = "A5816"
	a.balance = 150
	a.name = "Bob"
	a.typeA = "checking"

	fmt.Println(a)
	b := new(testaccount)
	b.number = "A4484"
	fmt.Println(b)

	c := testaccount{"A7156", 140, "Tim", "checking"}
	fmt.Println("Can't do values in different order", c)

	//Retrieve Struct Elements
	//Using Points
	p := &a
	(*p).balance = 240
	fmt.Println(a)
}

//go mod init mymodule
