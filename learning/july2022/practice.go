package july2022

import "fmt"

//import "math"
//import "strconv"

func Day1() {
	var name, address, year string = "Alex", "Ladson", "1991"

	fmt.Println(name, address, year)
	fmt.Println(&name)
	fmt.Println(&address)
	fmt.Println(&year)
}
