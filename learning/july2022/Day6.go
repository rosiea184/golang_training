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

//A struct within a struct

// package main

// import (
// 	"fmt"
// )

// type person struct {
// 	First string
// 	Last  string
// 	Age   int
// }

// type doubleZero struct {
// 	person
// 	First         string
// 	LicenseToKill bool
// }

// func main() {
// 	p1 := doubleZero{
// 		person: person{
// 			First: "James",
// 			Last:  "Bond",
// 			Age:   20,
// 		},

// 		First:         "Double Zero Seven",
// 		LicenseToKill: true,
// 	}

// 	p2 := doubleZero{
// 		person: person{
// 			First: "Miss",
// 			Last:  "MoneyPenny",
// 			Age:   19,
// 		},
// 		First:         "If looks could kill",
// 		LicenseToKill: false,
// 	}

// 	// fields and methods of the inner-type are promoted to the outer-type
// 	fmt.Println(p1.First, p1.person.First)
// 	fmt.Println(p2.First, p2.person.First)
// }
