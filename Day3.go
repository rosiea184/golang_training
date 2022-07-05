package main

import (
	"fmt"
	"strconv"
)

func main() {
	//print prime numbers
	// fmt.Print("3 ")
	// a := 4
	// b := 0

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
	// 	} else {
	// 		b = a + 2
	// 		for b == a+2 {
	// 			if b%2 == 0 {
	// 				b = 0
	// 			} else if b%3 == 0 {
	// 				b = 0
	// 			} else if b%5 == 0 && b != 5 {
	// 				b = 0
	// 			} else if b%7 == 0 && b != 7 {
	// 				b = 0
	// 			} else if b%11 == 0 && b != 11 {
	// 				b = 0
	// 			} else {
	// 				fmt.Print(a, " ", b, " ")
	// 				b = 0
	// 			}
	// 			a++
	// 		}
	// 	}
	// }
	// fmt.Println("")
	//Second assignment
	//a
	//ab
	//abc
	//abcd
	// beg := 'a'
	// end := 'd'
	// var r string
	// for i := beg; i <= end; i++ {
	// 	r = r + string(i)
	// 	fmt.Println(r)
	// }

	//looping through a string
	// var message string = "HELLO WORLD"
	// fmt.Println(message)

	// for i := 1; i < len(message); i = i + 2 {
	// 	fmt.Println(string(message[i]))
	// }
	// //The range function
	// for i, c := range message {
	// 	fmt.Print(i)
	// 	fmt.Println(string(c))
	// }
	//robert pike first name backwards only
	// name := "Robert Pike"
	// j := 0
	// for i := 0; i < len(name); i++ {
	// 	if string(name[i]) == " " {
	// 		j = i
	// 	}
	// }
	// for k := j - 1; k >= 0; k-- {
	// 	fmt.Print(string(name[k]))
	// }
	// for j < len(name) {
	// 	fmt.Print(string(name[j]))
	// 	j++
	// }
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
	//armstrong := 0
	var ret string
	//compare := 0
	//multi := 0
	//ones, tens, hundreds := 0, 0, 0
	for i := 1; i <= 10; i++ {
		ret = strconv.Itoa(i)
		for j := 0; j < len(ret); j++ {
			element, _ := strconv.Atoi(string((ret[j])))
			fmt.Println("The element is: ", element)
			// multi = element * element * element
			// fmt.Println(multi)
			// compare = compare + multi
			//fmt.Println("compare is: ", compare)
		}
		// if i == compare {
		// 	armstrong++
		// 	fmt.Println(i)
		// }
	}

	//Activity 22
	fmt.Println("Activity 22")
}
