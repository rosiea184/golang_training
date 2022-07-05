package main

import (
	"fmt"
	"strconv"
	//"math"
)

func main() {

	//activity 1
	// fmt.Print("Activity 1: ")
	// fmt.Println(((5 + 3) % 2) * 9)

	// //Activity 2
	// var floatNumber string
	// var floatInput float64
	// var intInput int

	// fmt.Print("Enter a float number: ")
	// fmt.Scan(&floatNumber)
	// floatInput, _ = strconv.ParseFloat(floatNumber, 64)
	// intInput = int(floatInput)
	// fmt.Println(intInput)

	//Activity 3
	// var P, r, n, t float64
	// var initialDeposit, interest, numYearInterst, numYear string
	// fmt.Print("Enter the initial deposit: ")
	// fmt.Scan(&initialDeposit)
	// fmt.Print("Enter the interest: ")
	// fmt.Scan(&interest)
	// fmt.Print("Enter the number of times per year interest is calculated: ")
	// fmt.Scan(&numYearInterst)
	// fmt.Print("Enter the the number of years since the initial deposit: ")
	// fmt.Scan(&numYear)
	// P, _ = strconv.ParseFloat(initialDeposit, 64)
	// r, _ = strconv.ParseFloat(interest, 64)
	// n, _ = strconv.ParseFloat(numYearInterst, 64)
	// t, _ = strconv.ParseFloat(numYear, 64)
	// var V = (P * math.Pow((1+r/n), (n*t)))
	// fmt.Println("The Value is: $", math.Round(V*100)/100)

	//Activity 4
	var princpleInt, rateInt, daysInt int
	var principle, rate, days string
	fmt.Print("Enter the initial deposit: ")
	fmt.Scan(&principle)
	fmt.Print("Enter the interest: ")
	fmt.Scan(&rate)
	fmt.Print("Enter the days: ")
	fmt.Scan(&days)
	princpleInt, _ = strconv.Atoi(principle)
	rateInt, _ = strconv.Atoi(rate)
	daysInt, _ = strconv.Atoi(days)
	var interest int = princpleInt * rateInt * daysInt / 365
	fmt.Println("The interest is: ", interest)

	//Activity 5
	// var a, b int = 0, 1
	// fmt.Println("a < b: ", a < b)
	// fmt.Println("a <= b: ", a <= b)
	// fmt.Println("a != b: ", a != b)
	// fmt.Println("a > b: ", a > b)
	// fmt.Println("a == b: ", a == b)
	// fmt.Println("a >= b: ", a >= b)

	//Activity 6
	// var value int
	// var input string
	// var inputFloat float32
	// var inputBin
	// fmt.Print("Choose a value")
	// fmt.Scan(&input)
	// value, _ = strconv.Atoi(input)
	// fmt.Println("You selected value: ", input)

	//Activity 7
	// var firstNumber, secondNumber, thirdNumber, fourthNumber, fifthNumber string
	// var firstInt, secondInt, thirdInt, fourthInt, fifthInt int
	// fmt.Print("Enter the first integer: ")
	// fmt.Scan(&firstNumber)
	// fmt.Print("Enter the second integer: ")
	// fmt.Scan(&secondNumber)
	// fmt.Print("Enter the third integer: ")
	// fmt.Scan(&thirdNumber)
	// fmt.Print("Enter the fourth integer: ")
	// fmt.Scan(&fourthNumber)
	// fmt.Print("Enter the fifth integer: ")
	// fmt.Scan(&fifthNumber)
	// firstInt, _ = strconv.Atoi(firstNumber)
	// secondInt, _ = strconv.Atoi(secondNumber)
	// thirdInt, _ = strconv.Atoi(thirdNumber)
	// fourthInt, _ = strconv.Atoi(fourthNumber)
	// fifthInt, _ = strconv.Atoi(fifthNumber)
	// var average int = ((firstInt + secondInt + thirdInt + fourthInt + fifthInt) / 5)

	// fmt.Print("The product: ")
	// fmt.Println(firstInt * secondInt * thirdInt * fourthInt * fifthInt)
	// fmt.Println("The average: ", average)
	// fmt.Print("The sum: ")
	// fmt.Println(firstInt + secondInt + thirdInt + fourthInt + fifthInt)

}
