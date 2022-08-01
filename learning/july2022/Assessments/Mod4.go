package assessments

import (
	"fmt"
	"strconv"
)

func PracticeForLoops() {
	//Activity 1
	// fmt.Println("Activity 1")
	// beg := 'A'
	// end := 'Z'
	// for i := beg; i <= end; i++ {
	// 	fmt.Print(string(i))
	// }
	// fmt.Println("")
	// //Activity 2
	// fmt.Println("Activity 2")
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Print(i, " ")
	// 	}
	// }
	// fmt.Println("")
	// for i := 1; i < 100; i++ {
	// 	if i%2 != 0 {
	// 		fmt.Print(i, " ")
	// 	}
	// }
	// fmt.Println("")
	// //Activity 3
	// fmt.Println("Activity 3")
	// fmt.Print("3 ")
	// a := 4
	// b := 2
	// for a <= 100 {
	// 	for b <= a {
	// 		if a%b == 0 && b != a {
	// 			b = a + 1
	// 		} else if b == a {
	// 			fmt.Print(a, " ")
	// 			b++
	// 		} else {
	// 			b++
	// 		}
	// 	}
	// 	a++
	// 	b = 2
	// }
	// for a <= 100 {
	// 	if a%2 == 0 {
	// 		a++
	// 	} else if a%3 == 0 {
	// 		a++
	// 	} else if a%5 == 0 && a != 5 {
	// 		a++
	// 	} else if a%7 == 0 && a != 7 {
	// 		a++
	// 	} else if a%11 == 0 && a != 11 {
	// 		a++
	// 	} else if a%13 == 0 && a != 13 {
	// 		a++
	// 	} else {
	// 		fmt.Print(a, " ")
	// 		a++
	// 	}
	// }
	fmt.Println("")
	// //Activity 4
	// fmt.Println("Activity 4")
	// factorial := 0
	// for i := 0; i <= 100; i++ {
	// 	factorial = factorial + i
	// }
	// fmt.Println(factorial)
	// //Activity 5
	// fmt.Println("Activity 5")
	// for i := 100; i <= 1000; i++ {
	// 	if i%50 == 0 {
	// 		fmt.Print(i, " ")
	// 	}
	// }
	// fmt.Println("")
	// // i:=100
	// // for i, range 1000{
	// // 	if i%50 == 0 {
	// // 		fmt.Print(i, " ")
	// // 	}
	// // }
	// //Activity 6
	// fmt.Println("Activity 6")
	// var userinput string
	// number := 0
	// fmt.Print("Enter a positive integer: ")
	// fmt.Scan(&userinput)
	// number, _ = strconv.Atoi(userinput)
	// if number <= 0 {
	// 	fmt.Println("Invalid Number")
	// } else {
	// 	for i := 1; i <= number; i++ {
	// 		output := number * i
	// 		fmt.Print(output, " ")
	// 	}
	// }
	// fmt.Println("")
	// //Activity 7
	// fmt.Println("Activity 7")
	// var base, exponent string
	// var x, y, result int
	// fmt.Print("Enter a base integer: ")
	// fmt.Scan(&base)
	// fmt.Print("Enter an exponent integer: ")
	// fmt.Scan(&exponent)
	// x, _ = strconv.Atoi(base)
	// y, _ = strconv.Atoi(exponent)
	// result = 1
	// for i := 1; i <= y; i++ {
	// 	result = result * x
	// }
	// fmt.Println(result)
	//Activity 8
	// fmt.Println("Activity 8")
	//Activity 9
	// fmt.Println("Activity 9")
	// var userint string
	// var digits, digit, sum, product, factorial int
	// fmt.Print("Enter an Integer: ")
	// fmt.Scan(&userint)
	// digits, _ = strconv.Atoi(userint)
	// fmt.Println("Number of digits: ", len(userint))
	// length := len(userint) - 1
	// fmt.Println("First Digit: ", string(userint[0]))
	// fmt.Println("Last Digit: ", string(userint[length]))
	// for c := range userint {
	// 	digit, _ = strconv.Atoi(string(userint[c]))
	// 	sum = sum + digit
	// }
	// fmt.Println("The sum is: ", sum)
	// digit = 0
	// product = 1
	// for c := range userint {
	// 	digit, _ = strconv.Atoi(string(userint[c]))
	// 	product = product * digit
	// }
	// fmt.Println("The product is: ", product)
	// for i := 2; i <= digits; i++ {
	// 	if digits%i == 0 && i != digits {
	// 		fmt.Println(digits, " is not a prime number")
	// 		break
	// 	} else if i == digits {
	// 		fmt.Println(digits, " is a prime number")
	// 	}
	// }
	// factorial = 1
	// for i := 1; i <= digits; i++ {
	// 	factorial = factorial * i
	// }
	// if factorial <= 0 {
	// 	fmt.Println("The number is too big to output correct factorial answer")
	// } else {
	// 	fmt.Println("The factorial of ", userint, " is ", factorial)
	// }
	// //Activity 10
	// fmt.Println("Activity 10")
	// var rows string
	// fmt.Print("Enter amount of rows: ")
	// fmt.Scan(&rows)
	// length, _ = strconv.Atoi(rows)
	// for i := 1; i <= length; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print(i)
	// 	}
	// 	fmt.Println("")
	// }
	// //Activity 11
	// fmt.Println("Activity 11")
	// var fst, snd string
	// var firstnum, secondnum, common int
	// fmt.Print("Enter 1st Number: ")
	// fmt.Scan(&fst)
	// fmt.Print("Enter 2nd Number: ")
	// fmt.Scan(&snd)
	// firstnum, _ = strconv.Atoi(fst)
	// secondnum, _ = strconv.Atoi(snd)

	// for i := 2; i < firstnum; i++ {
	// 	if firstnum%i == 0 {
	// 		if secondnum%i == 0 {
	// 			common = i
	// 		}
	// 	}
	// }
	// fmt.Println("The Greatest Common Divisor is ", common)
	// //Activity 12
	// fmt.Println("Activity 12")
	// for i := firstnum; i > 2; i-- {
	// 	if firstnum%i == 0 {
	// 		if secondnum%i == 0 {
	// 			common = i
	// 		}
	// 	}
	// }
	// fmt.Println("The Lowest Common Divisor is ", common)
	// //Activity 13
	// fmt.Println("Activity 13")
	// //Activity 14
	// fmt.Println("Activity 14")
	// reverse := ""
	// fmt.Print("Enter a two diget or more number: ")
	// fmt.Scan(&reverse)
	// for i := len(reverse) - 1; i >= 0; i-- {
	// 	fmt.Print(string(reverse[i]))
	// }
	// fmt.Println("")
	//Activity 15
	// fmt.Println("Activity 15")
	// act15 := ""
	// fmt.Print("Enter a string: ")
	// fmt.Scan(&act15)
	// for i := 0; i < len(act15); i++ {
	// 	fmt.Println(string(act15[i]))
	// }
	// //Activity 16
	// fmt.Println("Activity 16")
	// counter := 0
	// for i := 0; i < len(act15); i++ {
	// 	counter++
	// }
	// fmt.Println("The Length of the String is: ", counter)
	//Activity 21
	fmt.Println("Activity 21")
	armstrong := 0
	var ret string
	compare := 0
	multi := 0
	//ones, tens, hundreds := 0, 0, 0
	for i := 1; i <= 10; i++ {
		ret = strconv.Itoa(i)
		for j := 0; j < len(ret); j++ {
			element, _ := strconv.Atoi(string((ret[j])))
			fmt.Println("The element is: ", element)
			multi = element * element * element
			fmt.Println(multi)
			compare = compare + multi
			fmt.Println("compare is: ", compare)
		}
		if i == compare {
			armstrong++
			fmt.Println("armstong is: ", i)
			compare = 0
		} else {
			compare = 0
		}
	}

	//Activity 22
	/*
			Write a program that generates a random integer between 0 and 10 and asks the user to guess what the number is.

		    If the user's guess is higher than the random number, the program should display "Too high, try again."
		    If the user's guess is lower than the random number, the program should display "Too low, try again."
		    If the user's number is the same as the random number, the program should display, "That's right. You guessed the number!"

		The program should use a loop that repeats until the user correctly guesses the random number.
	*/
	fmt.Println("Activity 22")
}
func CoffeeShop() {
	/*
		Write a program that will calculate the cost of a custom cup of coffee at a gourmet coffee shop, based on the size of the cup, the type of coffee selected, and flavors that can be added to the coffee. It should complete the following steps:

			Ask the user what size cup they want, choosing between small, medium, and large.
			Ask the user what kind of coffee they want, choosing between brewed, espresso, and cold brew.
			Ask the user what flavoring they want, if any. Choices include hazelnut, vanilla, and caramel.
			Calculate the price of the cup using the following values:
			    Size:
			        small: $2
			        medium: $3
			        large: $4
			    Type:
			        brewed: no additional cost
			        espresso: 50 cents
			        cold brew: $1
			    Flavoring:
			        None: no additional cost
			        All other options: 50 cents
			Display a statement that summarizes what the user ordered.
			Display the total cost of the cup of coffee as well as the cost with a 15% tip, in phrases that explain the values to the user. Round the cost with tip to two decimal places.
			    For example, if the user asks for a medium-sized espresso with hazelnut flavoring, the total should be $4; the total with a tip should be $4.60.

		Sample
		Do you want small, medium, or large? small
		Do you want brewed, espresso, or cold press? espresso
		Do you want a flavored syrup? (Yes or No) yes
		Do you want hazelnut, vanilla, or caramel? vanilla
		You asked for a small cup of espresso coffee with vanilla syrup.
		Your cup of coffee costs 3.0
		The price with a tip is 3.45

		Requirements

			Add your name and a current date to a commented line at the top of the code.
			Use meaningful variable names to clearly identify the values in the program.
			The user should be able to enter text values using any case without causing errors.
			All output should be clear and meaningful to the user.

		Checklist

			When you have finished, use the following checklist to make sure you have completed all parts of the task:

			    Your name and a current date appear in a comment at the top of the program.
			    The program prompts the user for the size of the coffee and accepts the values small, medium, or large as input.
			    The program prompts the user for the type of coffee and accepts the values brewed, espresso, or cold brew as input.
			    The program prompts the user for a flavoring and accepts the values none, hazelnut, vanilla, or caramel.
			    The program displays a summary of the user's order.
			    The program displays a total that corresponds to the price of the cup of coffee based on the user's choices.
			    The program displays a total that corresponds to the price of the cup of coffee with a 15% tip.

	*/

}

func FizzBuzz() {
	/*
		In this assessment, you will write a program that loops through a series of values and uses those values to determine the output shown to the user.
		Instructions

		Write a program that performs the following steps:

		    Ask the user for a number.
		    Output a count starting with 0.
		        Display the count number if it is not divisible by 3 or 5.
		        Replace every multiple of 3 with the word "fizz."
		        Replace every multiple of 5 with the word "buzz."
		        Replace multiples of both 3 and 5 with "fizz buzz."
		    Continue counting until the number of integers replaced with "fizz," "buzz," or "fizz buzz" reaches the input number.
		    The last output line should read "TRADITION!!"

		For example:

		How many fizzing and buzzing units do you need in your life? 7
		0
		1
		2
		fizz
		4
		buzz
		fizz
		7
		8
		fizz
		buzz
		11
		fizz
		13
		14
		fizz buzz
		TRADITION!!

		Requirements

		    Add your name and the current date as a comment on the first line of code.
		    Use meaningful names for all variables so that another developer can understand the code easily.
		    All text displayed to the user must be meaningful and use correct grammar and spelling.
		    The program should prompt the user for a number.
		    The program should output a count that replaces individual values with "fizz," "buzz," or "fizz buzz," as indicated in the instructions.
		    The program should output the text "TRADITION!!" as the last line of output, as indicated in the instructions.
	*/
}
