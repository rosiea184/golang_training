package july2022

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

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
func activityOneTwo() {
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
}
func activityThree() {
	var n1, n2, n3, n4, n5 int
	fmt.Print("Please give 5 numbers: ")
	fmt.Scan(&n1, &n2, &n3, &n4, &n5)
	sum := n1 + n2 + n3 + n4 + n5
	product := n1 * n2 * n3 * n4 * n5
	fmt.Println("The numbers you entered are: ", n1, n2, n3, n4, n5)
	fmt.Println("The Sum of those numbers are, ", sum, " and the product is, ", product)
}

type wordStorage struct {
	word   string
	amount int
}

func activityFour() {
	//wordcounter := 0
	distinctWords := 0
	paragraph := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	paragraph = strings.ToLower(paragraph)
	wordString := strings.Split(paragraph, " ")
	fmt.Println(wordString)
	fmt.Println(len(wordString))
	for i := 0; i < len(wordString)-1; i++ {
		for j := i + 1; j < len(wordString); j++ {
			if wordString[i] != wordString[j] {
				if j == len(wordString)-1 {
					//Determine the number of distinct words in the lowercase string.
					distinctWords++
					//fmt.Println("I made it inside: ", distinctWords)
				}
			} else {
				break
			}
		}
	}
	// Activity 4

	// In this activity, we will start with a paragraph of text and end up with a frequency count of the distinct words in that text.

	// Start with a paragraph of your choice. You may use any source for the paragraph.

	// After inputting the paragraph, perform the following operations on it:

	//     Calculate the number of times each word appears in the lowercase string.
	//         Use the appropriate data structure to store the words and their frequency of occurrence.
	//     Remove the punctuation from the lowercase string.
	//     Perform a count analysis on the text without punctuation characters.
}
func createSlice() []int {
	//     a function that creates a slice with a collection of random numbers between 0 and 100
	n := rand.Intn(10) + 5
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}
func sortSlice(x []int) []int {
	asc := make([]int, len(x))
	index := 0
	length := len(x)
	var place []int
	for i := 0; i < len(x); i++ {
		place = append(place, x[i])
	}
	for i := 0; i < length; i++ {
		asc[i], index = findMin(place)
		place = pop2(place, index)
	}
	return asc
	//     a function that sorts a slice in-place (ascending and descending depending on an input parameter)
}
func combineSlice(x []int, y []int) []int {
	//     a function that takes as input two sorted slices and combines the slices into a single sorted slice.
	var z []int
	for i := 0; i < len(x); i++ {
		z = append(z, x[i])
	}
	for i := 0; i < len(y); i++ {
		z = append(z, y[i])
	}
	z = sortSlice(z)
	return z
}
func activityFive() {
	rand.Seed(time.Now().UnixNano())
	x := createSlice()
	y := createSlice()
	fmt.Println("The created slices")
	fmt.Println(x)
	fmt.Println(y)
	x = sortSlice(x)
	y = sortSlice(y)
	fmt.Println("The sorted slices")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("The combined slices")
	fmt.Println(combineSlice(x, y))
}

type cube struct {
	length float64
}
type box struct {
	length, width, height float64
}
type sphere struct {
	radius float64
}

func (s cube) volume() float64 {
	v := s.length * s.length * s.length
	return v
}
func (s box) volume() float64 {
	v := s.length * s.width * s.height
	return v
}
func (s sphere) volume() float64 {
	v := (4 / 3) * math.Pi * (s.radius * s.radius * s.radius)
	return v
}

func activitySix() {
	// Create a program that will calculate the volume of a solid.
	sphere := sphere{}
	cube := cube{}
	box := box{5, 3, 4}
	cube.length = 5
	sphere.radius = 3
	fmt.Println(sphere.volume())
	fmt.Println(cube.volume())
	fmt.Println(box.volume())
}

func Mod6DataCollections() {
	//Data Collections
	//activityOneTwo()
	//activityThree
	//activityFour()
	//activityFive()
	activitySix()
}

// Challenge

// After verifying that the program meets the requirements outlined above and works correctly, add additional shapes. For example, you could add a cuboid, a cone, and a pyramid.
