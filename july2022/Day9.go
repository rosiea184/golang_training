package july2022

import (
	"fmt"
	"math/rand"
	"time"
)

type Polygons interface {
	Peremeter()
}
type Object interface {
	NumberOfSide()
}
type Pentagon int

func (p Pentagon) Peremeter() {
	fmt.Println("Pentagon", 5*p)
}
func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}
func shapes() {
	var p Pentagon = Pentagon(50)
	p.Peremeter()
	var obj Object = Pentagon(20)
	obj.NumberOfSide()
}

//Activities
func activity1Random(n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(200) - 100
	}
	return arr
}
func findMax(arr []int) (int, int) {
	max := -101
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}
	return max, index
}
func findMin(arr []int) (int, int) {
	min := 101
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			index = i
		}
	}

	return min, index
}

func pop2(place []int, x int) []int {
	return append(place[:x], place[x+1:]...)[:len(place)-1]
}
func ascending(arr []int) []int {
	asc := make([]int, len(arr))
	index := 0
	length := len(arr)
	var place []int
	for i := 0; i < len(arr); i++ {
		place = append(place, arr[i])
	}
	for i := 0; i < length; i++ {
		asc[i], index = findMin(place)
		place = pop2(place, index)
	}
	return asc
}

func descending(arr []int) []int {
	des := make([]int, len(arr))
	index := 0
	length := len(arr)
	var place []int
	for i := 0; i < len(arr); i++ {
		place = append(place, arr[i])
	}
	for i := 0; i < length; i++ {
		des[i], index = findMax(place)
		place = pop2(place, index)
	}
	return des
}
func mean(arr []int) int {
	mean := 0
	for i := 0; i < len(arr); i++ {
		mean += arr[i]
	}
	mean = mean / len(arr)
	return mean
}
func median(arr []int) int {
	half := len(arr) / 2
	median := arr[half]
	return median
}
func postiveArray(arr []int) []int {
	var place []int
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			place = append(place, arr[i])
		}
	}
	return place
}
func negativeArray(arr []int) []int {
	var place []int
	for i := 0; i < len(arr); i++ {
		if arr[i] <= 0 {
			place = append(place, arr[i])
		}
	}
	return place
}
func longSequence(arr []int) []int {
	var place []int
	for i := 0; i < len(arr); i++ {
		place = append(place, arr[i])
	}
	j := 0
	length := 1
	count := 0
	var output []int
	for i := 1; i < len(arr); i++ {
		if place[i] >= place[i-1] {
			length++
		} else if place[i] < place[i-1] {
			if length > count {
				count = length
				output = append(place[j:i])
				j = i
				length = 1
			} else {
				length = 1
				j = i
			}
		}
	}
	return output
}
func removeDup(arr []int) []int {
	var place []int
	//length := len(arr)
	for i := 0; i < len(arr); i++ {
		place = append(place, arr[i])
	}
	for i := 0; i < len(place); i++ {
		for j := 1; j < len(place); j++ {
			if place[i] == place[j] {
				place = pop2(place, j)
			}
		}
	}
	return place
}

func Day9() {
	//Data Collections
	arr := activity1Random(20)
	fmt.Println(arr)
	max, index1 := findMax(arr)
	min, index2 := findMin(arr)
	fmt.Println("The largest number is: ", max, " at index: ", index1)
	fmt.Println("The smallest number is: ", min, " at index: ", index2)
	dec := descending(arr)
	asc := ascending(arr)
	fmt.Println("Ascending order is: ", asc)
	fmt.Println("Decending order is: ", dec)
	fmt.Println("The mean is: ", mean(arr))
	fmt.Println("The median is: ", median(arr))
	fmt.Println("All the positive integers is: ", postiveArray(arr))
	fmt.Println("All the negative integers is: ", negativeArray(arr))
	fmt.Println("Longest sequence of sorted numbers is: ", longSequence(arr))
	//fmt.Println("The array without duplicates is: ", removeDup(arr))
	// Activity 3
}

// Create a program that completes the following tasks:

//     It prompts the user for a series of 5-10 integers.
//         The user must be prompted for a minimum of 5 numbers.
//         After the user has entered at least 5 numbers, the program should allow the user to stop entry or enter another number.
//     When the user chooses to stop or after 10 numbers have been entered, the program stops prompting for values and performs the following calculations:
//         the product of the integers
//         the sum of the integers
//     After performing the calculations, the program should display the following to the user:
//         the values the user entered
//         each of the calculations, using a phrase that identifies the value

// Use the appropriate data structure to store the numbers.
// Activity 4

// In this activity, we will start with a paragraph of text and end up with a frequency count of the distinct words in that text.

// Start with a paragraph of your choice. You may use any source for the paragraph.

// After inputting the paragraph, perform the following operations on it:

//     Convert all text to lowercase characters.
//     Split the lowercase string into individual words.
//     Count the number of words in the lowercase string.
//     Determine the number of distinct words in the lowercase string.
//     Calculate the number of times each word appears in the lowercase string.
//         Use the appropriate data structure to store the words and their frequency of occurrence.
//     Remove the punctuation from the lowercase string.
//     Perform a count analysis on the text without punctuation characters.

// Activity 5

// Create a program that includes the following features:

//     a function that creates a slice with a collection of random numbers between 0 and 100
//     a function that sorts a slice in-place (ascending and descending depending on an input parameter)
//     a function that takes as input two sorted slices and combines the slices into a single sorted slice.
//     In the main function
//         create two slices
//         sort the two slices
//         combine and sort the two slices

// Additional requirements:

//     Display all slices to the user, with appropriate feedback to identify each slice.
//     Do not use any built-in functions

// Activity 6

// Create a program that will calculate the volume of a solid.

// Start with the following structs:

//     Cube: represents a cube with one attribute: length: float64
//     Box: represents a box with three attributes: length, width, height as float64
//     Sphere represents a sphere with one attribute: radius: float64

// Implement a volume method for each of the structs defined above. The volume method returns the volume of a cube, or box, or sphere.

// Use the main function to create different shapes and compute their volume.
// Challenge

// After verifying that the program meets the requirements outlined above and works correctly, add additional shapes. For example, you could add a cuboid, a cone, and a pyramid.
