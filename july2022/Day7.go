package july2022

import (
	"fmt"
	//"golang.org/x/tour/pic"
)

func SliceActivityString() {
	//Slice Activity Strings
	word1 := ""
	word2 := ""
	word3 := ""
	word4 := ""
	word5 := ""
	word6 := ""
	word7 := ""
	word8 := ""
	word9 := ""
	word10 := ""
	fmt.Print("Imput ten words: ")
	fmt.Scan(&word1, &word2, &word3, &word4, &word5, &word6, &word7, &word8, &word9, &word10)
	activity := []string{word1, word2, word3, word4, word5, word6, word7, word8, word9, word10}
	average := 0
	for i := 0; i < len(activity); i++ {
		average = +len(activity[i])
	}
	average = average / len(activity)
	var short []string
	var long []string
	for i := 0; i < len(activity); i++ {
		if len(activity[i]) < average {
			short = append(short, string(activity[i]))
		} else if len(activity[i]) >= average {
			long = append(long, string(activity[i]))
		}
	}
	fmt.Println("These are the words you gave: ", activity)
	fmt.Println("These words are shorter than the average: ", short)
	fmt.Println("These words are longer than the average: ", long)
}

type class struct {
	ClassName string
	students  []student //student is user defined data type
}
type student struct {
	Name   string
	RollNo int
	City   string
}

func changecolumn(matrix [][]int) {
	/*
		input
		1 2 3
		4 5 6
		7 8 9
		output(interchange the first and last column)
		3 2 1
		6 5 4
		9 8 7
	*/
	var temp int
	for i := 0; i < len(matrix); i++ {
		temp = matrix[i][0]
		matrix[i][0] = matrix[i][2]
		matrix[i][2] = temp
	}
	for _, ind := range matrix {
		fmt.Println(ind)
	}
}
func changerow(matrix [][]int) {
	/*
		input
		1 2 3
		4 5 6
		7 8 9
		output(interchange the first and last row)
		7 8 9
		4 5 6
		1 2 3
	*/
	var temp2 []int
	for j := 0; j < len(matrix); j++ {
		temp2 = append(temp2, matrix[0][j])
	}
	matrix[0] = matrix[2]
	matrix[2] = temp2
	for _, ind := range matrix {
		fmt.Println(ind)
	}

}

// func Pic(dx, dy int) [][]uint8 {
// 	var picture [][]uint8
// }

func Day7() {
	//New Structure
	abs := student{Name: "Ross", RollNo: 30, City: "NewYork"}
	cbs := student{Name: "Marry", RollNo: 31, City: "London"}
	students := []student{abs, cbs, student{Name: "Jack", RollNo: 32, City: "London"}}
	//fmt.Println(students)
	students = append(students, student{Name: "Kate", RollNo: 33, City: "NewYork"})
	//slice of slice of structure
	class := class{"firstA", students}
	fmt.Println(class)
	for i, j := range students {
		fmt.Println(i, "   ", j)
	}
	//Slices
	//pointer to an array/subset
	//Len, Cap
	s := make([]int, 1, 1)
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d , %d ", len(s), cap(s))
		s = append(s, i)
	}
	var records [][]string
	students1 := make([]string, 3)
	students1[0] = "ABC"
	students1[1] = "XYZ"
	students1[2] = "AXZ"

	records = append(records, students1)
	fmt.Println(records)
	students2 := make([]string, 3)
	students2[0] = "123"
	students2[1] = "456"
	students2[2] = "789"

	records = append(records, students2)
	fmt.Println(records)
	//start of assignment
	var matrix [][]int
	row1 := make([]int, 3)
	row1[0] = 1
	row1[1] = 2
	row1[2] = 3

	matrix = append(matrix, row1)
	fmt.Println(matrix)
	row2 := make([]int, 3)
	row2[0] = 4
	row2[1] = 5
	row2[2] = 6

	matrix = append(matrix, row2)
	fmt.Println(matrix)

	row3 := make([]int, 3)
	row3[0] = 7
	row3[1] = 8
	row3[2] = 9

	matrix = append(matrix, row3)
	fmt.Println(matrix)
	for _, ind := range matrix {
		fmt.Println(ind)
	}
	fmt.Println(matrix[0][0])
	fmt.Println(matrix[1][0], matrix[1][1])
	fmt.Println(matrix[2][0], matrix[2][1], matrix[2][2])

	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println("")
	}
	fmt.Println("")
	changecolumn(matrix)
	fmt.Println("")
	changecolumn(matrix)
	fmt.Println("")
	changerow(matrix)

	//Slices Activity: Bitmap Images
	//pic.Show(Pic)

}
