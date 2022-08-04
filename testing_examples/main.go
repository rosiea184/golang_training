package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(getArr())

	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}
	name, exists := getName(os.Args[1])
	if !exists {
		fmt.Printf("error: user (%v) not found", os.Args[1])
		os.Exit(1)
	}
	fmt.Println("Hi,", name)

}

func getArr() [10]int {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = 1
	}
	return arr
}

var users = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

func getName(id string) (string, bool) {
	name, exists := users[id]
	return name, exists
}
