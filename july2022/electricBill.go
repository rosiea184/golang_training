package july2022

import (
	"fmt"
	"strconv"
)

func ElectricBill() {
	var rate1, rate2, rate3 int = 5, 7, 10
	var units int
	var unitsInput string
	var output int

	fmt.Print("Enter Number of Units: ")
	fmt.Scan(&unitsInput)
	units, _ = strconv.Atoi(unitsInput)
	if units <= 100 {
		output = units * rate1
		fmt.Println("Your bill is: ", output)
	} else if units > 100 && units <= 200 {
		output = units * rate2
		fmt.Println("Your bill is: ", output)
	} else if units > 200 && units <= 350 {
		output = units * rate3
		fmt.Println("Your bill is: ", output)
	} else {
		fmt.Println("Invalid Input")
	}

}
