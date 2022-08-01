package assessments

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Calculator() {
	/*
		Create a calculator application that runs from the command line.

		The calculator should perform the following tasks:

		    The user can enter a word like quit or exit at any prompt to end the program.
		    Accept two numeric values from the user.
		    Allow the user to choose a mathematical operation using the two input values, including at least addition, subtraction, multiplication, modulus (%), square root, and factorial. You may include other operations if you wish.
		    Display the output as a complete mathematical statement. For example, if the user wants to add two numbers, the output should look like: 8 + 9 = 17
		    Prompt the user to start over after completing a calculation.

		Additional requirements:

		    Use and call appropriate functions in the code.
		        Each operation should be a separate function.
		        Include additional functions as appropriate.
		    If the user enters an unexpected/invalid value at a prompt or the output causes an error (such as divide-by-0), the program should display appropriate feedback and prompt the user to try again.
		    The program should be as user-friendly as possible, by giving the user clear options for each prompt and letting them know how to exit the program when they wish.
	*/
}
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

func Pic(dx, dy int) [][]uint8 {
	//var picture [][]uint8
}
func SliceBitmapImage() {
	pic.Show(Pic)
}

func MapsKeywordSearch() {
	/*
		Write a program that includes the following:

		    Create a map that includes at least ten key:value pairs.
		        The pairs should make sense, but you can choose whatever you wish.
		        Suggestions include produce categories (e.g., apple:fruit and onion:vegetable), animal categories (e.g., mammals, birds, fish), city populations, or word definitions.
		    Prompt the user to input a search term.
		    Display all key:value pairs that include the input search term, either in the key or in the value.
		        For example, if the topic is produce, the user can enter either apple or fruit and see all matching map entries.
		    If the value does not exist in the dictionary, display a user-friendly error message.
		    Prompt the user to start over again or exit the program.
	*/
	//category := make(map[string]string)
	//category[""]=""
	//category[""]=""
	//category[""]=""
	//category[""]=""
	//category[""]=""
	//category[""]=""
	//category[""]=""
	//category[""]=""
	//category[""]=""
	//category[""]=""
}
