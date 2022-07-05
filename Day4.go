package main

import (
	"fmt"
	"strconv"
)

func fibonacii() func() int {
	first, second := 0, 1
	return func() int {
		temp := first
		first, second = second, first+second
		return temp
	}
}

func arrayAssign() {
	// 1. position value
	// 2. direction (left or right)
	//Program should shift the entire array either left or right
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//right{9,1,2,3,4,5,6,7,8} left{2,3,4,5,6,7,8,9,1}
	posS, dirS := "", ""
	pos := 0
	dir := 0
	fmt.Print("Please pick a position between 0-9: ")
	fmt.Scan(&posS)
	fmt.Print("Please pick a direction, 1 for right, 2 for left: ")
	fmt.Scan(&dirS)
	pos, _ = strconv.Atoi(string(posS))
	dir, _ = strconv.Atoi(string(dirS))
	if dir == 1 {
		length := len(arr) - (pos)
		arr = append(arr[length:], arr[0:length]...)
		fmt.Println(arr)
	} else if dir == 2 {
		arr = append(arr[pos:], arr[0:pos]...)
		fmt.Println(arr)
	} else {
		fmt.Println("Invalid direction number")
	}

}

func main() {
	//fibonachii
	f := fibonacii()
	for i := 0; i <= 5; i++ {
		fmt.Println(f())
	}
	//arrayAssign()

	//Array Practice
	var numbers [3]int

	numbers[0] = 1
	numbers[2] = 3455
	numbers[1] = 34
	//numbers[3] = 11

	fmt.Println(numbers)

	//Q1 No it does not need to be sequential
	//Q2 It leaves the empty one 0 if you do not put a value in it
	//Q3 It will replace the previous value
	//Q4 Out of bounds error, will not complie
	//Q5 Will not complie

	var first [10]int
	var second [5]string

	first[0] = 1
	first[1] = 2
	first[2] = 3
	first[3] = 4
	first[4] = 5
	first[5] = 6
	first[6] = 7
	first[7] = 8
	first[8] = 9
	first[9] = 10

	second[0] = "H"
	second[1] = "E"
	second[2] = "L"
	second[3] = "L"
	second[4] = "O"
	fmt.Println(first)
	fmt.Println(second)

	var first2 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	second2 := [5]string{"H", "E", "L", "L", "O"}

	fmt.Println(first2)
	fmt.Println(second2)

	//Array Size

	//Two Dimensional Arrays
	fmt.Println("Two Dimentional Arrays")
	matrix := [4][4]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
		{7, 8, 9, 10},
		{10, 11, 12, 13},
	}
	fmt.Println(matrix, len(matrix))
	fmt.Println(matrix[1], len(matrix[1]))
	fmt.Println(matrix[3][3])
}
