package july2022

import (
	"fmt"
	"regexp"

	//"sort"
	"strconv"
	"time"
)

// create an alias type
type mytype []string

// implement the Len method
func (s mytype) Len() int {
	return len(s)
}

// implement the Swap method
func (s mytype) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// implement the Less method
func (s mytype) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func Day16() {
	/* Practice Stuff
	// define a slice
	numbers := []int{67, 18, 62, 60, 25, 64, 75, 5, 17, 55}
	fmt.Println("Original Numbers:", numbers)

	// use the sort.Ints function to sort the values in the slice
	sort.Ints(numbers)

	fmt.Println("Sorted Numbers:", numbers)
	// define a slice
	words := []string{"camel", "zebra", "horse", "dog", "elephant", "giraffe"}
	fmt.Println("Original slice:", words)
	fmt.Println("The original values are sorted:", sort.StringsAreSorted(words))

	// sort the values in the slice
	sort.Strings(words)
	fmt.Println("Sorted slice:", words)
	fmt.Println("The values are sorted:", sort.StringsAreSorted(words))
	// create a slice of strings
	fruits := []string{"pear", "passionfruit", "mango", "banana", "fig"}
	fmt.Println("Original slice:", fruits)

	// create a mytype variable
	myfruits := mytype(fruits)

	sort.Sort(myfruits)
	fmt.Println("Sorted by length:", myfruits)
	//--------------------
	// display current time
	now := time.Now()
	fmt.Println("Today's date and time:", now)
	fmt.Println("Current year:", now.Year())
	fmt.Println("Current month:", now.Month())
	fmt.Println("Current day:", now.Day())
	fmt.Println("Current hour:", now.Hour())
	fmt.Println("Current minute:", now.Minute())
	fmt.Println("Current second:", now.Second())
	fmt.Println("Current nanosecond:", now.Nanosecond())
	fmt.Println("Current location:", now.Location())
	fmt.Println("Current weekday:", now.Weekday())
	fmt.Println(now.Format(time.RFC1123Z))
	t1, e := time.Parse(
		time.RFC3339,
		"2025-05-21T12:50:41+00:00")

	fmt.Println(t1)
	fmt.Println(t1.Day())
	fmt.Println(t1.Month())

	// error if there is an error during parsing;
	// nil if there is no error
	fmt.Println(e)
	//-------------------------------
	// check if string starts with C and ends with n
	m1, err1 := regexp.MatchString("C([a-z]+)n", "Catelyn")
	fmt.Println(m1)
	fmt.Println(err1)

	// check if string contains at least one digit
	m2, err2 := regexp.MatchString("[0-9]", "jonathan6smith")
	fmt.Println(m2)
	fmt.Println(err2)
	*/
	//email
	// m2 := regexp.MustCompile("[0-9]")
	// fmt.Println(m2.MatchString("jonathan6smith"))
	email := regexp.MustCompile(`^([\w-\.]+)@([\w-]+\.)+[\w-]{2,4}$`)
	fmt.Println(email.MatchString("joe.j@aol.co.uk"))
	number := regexp.MustCompile(`^[0-9]+[\(\-]([0-9]{3})+[\)\-]([0-9]{3})+\-([0-9]{4})$`)
	fmt.Println(number.MatchString("1(843)666-6825"))
	ipAdd := regexp.MustCompile(`[192]\.1([67]{1})+(\d{1})\.\d{1}`)
	fmt.Println(ipAdd.MatchString("192.175.1"))
	//Find any three-letter words that start with the same letter and end with the same letter, but which might have a different letter in between, such as cat or cot.

	//threeWord := regexp.MustCompile(`([a-z][a-z][a-z]){3}`)

	//example
	/*
			func mymatch(name string){
		    x:=[]byte(name)

		    regEx1,_:=regexp.Compile("^"+string(x[0])+".*"+string(x[0])+"$")

		    match1:=regEx1.Match([]byte(name))
		    fmt.Println(match1)
		}

		func main(){
		       mymatch("abbba")
		       mymatch("cbc")
		       mymatch("cdddc")
		       mymatch("adc")    }

	*/

	// //given
	// var ip string
	// str := "This is a string contain IP  and yo need to parse  255.254.1.2    and 192.168.1.1"

	// ip = findip(str)
	// ip2 := findip2(str)
	// fmt.Println(ip, " ", ip2)

}

//255.254.0.1
func findip(input string) string {
	pt := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])" //"(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9]])"
	regExParretn := pt + "\\." + pt + "\\." + pt + "\\." + pt
	regEx := regexp.MustCompile(regExParretn)
	fmt.Println(regEx.String())
	return regEx.FindString(input)
}

//192.168.1.1
func findip2(input string) string {
	pt := "(192|1[6-7][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regExParretn := pt + "\\." + pt + "\\." + pt + "\\." + pt
	regEx := regexp.MustCompile(regExParretn)
	fmt.Println(regEx.String())
	return regEx.FindString(input)
}

func leapYear(y string) bool {
	leapYear := false
	uy, _ := strconv.Atoi(string(y))
	if uy%4 == 0 {
		leapYear = true
	} else {
		leapYear = false
	}
	return leapYear
}
func DateTimeAct() {
	//Activity 1
	fmt.Println("Activity 1")
	now := time.Now()
	fmt.Println("Today's date and time:", now)
	fmt.Println("Current year:", now.Year())
	fmt.Println("Current month:", now.Month())
	_, wk_num := now.ISOWeek()
	fmt.Println("Current week:", wk_num)
	fmt.Println("Current day:", now.Day())
	fmt.Println("Current day of year:", now.YearDay())
	fmt.Println("Current weekday:", now.Weekday())

	//Activity 2
	fmt.Println("Activity 2")
	userYear := ""
	fmt.Print("Enter a year to see if it is a leap year: ")
	fmt.Scan(&userYear)
	fmt.Println("Your year ", userYear, " is: ", leapYear(userYear))
	//Activity 3
	// 	Write a function that prompts the user for a date (or uses the current date if the user does not input a date) and subtracts five days from that date. The function should return the new date.
	// Display the result to the user.
	fmt.Println("Activity 3")
	diff := now.AddDate(0, 0, -5)
	fmt.Println("Five days ago:", diff.Format(time.RFC1123Z))
	//Activity 4
	//Write a program that converts a Unix timestamp string to a human-readable date.
	fmt.Println("Activity 4")
	unixtime := now.Unix()
	timeU := time.Unix(unixtime, 0)
	fmt.Println(timeU.Format(time.RFC1123Z))
	//Activity 5
	//Write a program that prints the dates for yesterday, today, and tomorrow.
	fmt.Println("Activity 5")
	yesterday := now.AddDate(0, -1, 0)
	tomorrow := now.AddDate(0, 1, 0)
	fmt.Println("Yesterday: ", yesterday.Format(time.RFC1123Z))
	fmt.Println("Today: ", now.Format(time.RFC1123Z))
	fmt.Println("Tomorrow: ", tomorrow.Format(time.RFC1123Z))
	//Activity 6
	//Write a program that prints the date for the next five days, starting from today.
	fmt.Println("Activity 6")
	diff5 := now.AddDate(0, 0, 5)
	fmt.Println("Five days later:", diff5.Format(time.RFC1123Z))
	//Activity 7
	//Write a program that adds five seconds to the current time and displays the result.
	fmt.Println("Activity 7")
	diff5Sec := now.Add(5 * time.Second)
	fmt.Println("Five seconds later:", diff5Sec.Format(time.RFC1123Z))
	//Activity 8
	//Write a program that prompts the user for two different dates, computes the number of days between those dates, and displays the results.
	//The order in which the user enters the dates should not affect the results.
	fmt.Println("Activity 8")
	uDate1Y, uDate1M, uDate1D, uDate2Y, uDate2M, uDate2D := "", "", "", "", "", ""
	fmt.Print("Enter first date by Year Month Day in number format: ")
	fmt.Scan(&uDate1Y, &uDate1M, &uDate1D)
	fmt.Print("Enter second date by Year Month Day in number format: ")
	fmt.Scan(&uDate2Y, &uDate2M, &uDate2D)
	u1y, _ := strconv.Atoi(string(uDate1Y))
	u1m, _ := strconv.Atoi(string(uDate1M))
	u1d, _ := strconv.Atoi(string(uDate1D))
	u2y, _ := strconv.Atoi(string(uDate2Y))
	u2m, _ := strconv.Atoi(string(uDate2M))
	u2d, _ := strconv.Atoi(string(uDate2D))
	date1 := time.Date(u1y, time.Month(u1m), u1d, 00, 00, 0, 0, time.Local)
	date2 := time.Date(u2y, time.Month(u2m), u2d, 00, 00, 0, 0, time.Local)
	redate := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	match1 := redate.FindString(date1.String())
	match2 := redate.FindString(date2.String())
	if date1.Before(date2) == true {
		diff := date2.Sub(date1)
		days := diff.Hours() / 24
		fmt.Println("The number of days between ", match1, " and ", match2, " is: ", days)
	} else if date1.After(date2) == true {
		diff := date1.Sub(date2)
		days := diff.Hours() / 24
		fmt.Println("The number of days between ", match2, " and ", match1, " is: ", days)
	}
}
func RegularExpAct() {
	fmt.Println("Activity 1")
	//Write a program to check that a string contains only letters and numbers (e.g., a-z, A-Z, 0-9).
	act1 := regexp.MustCompile(`[a-zA-Z0-9]*`)
	fmt.Println(act1.FindString("yes"))
	fmt.Println("Activity 2")
	//Write a program that finds all strings that include the letter i followed by zero or more instances of the letter n.
	act2 := regexp.MustCompile(`in*`)
	fmt.Println(act2.FindString("input"))
	fmt.Println("Activity 3")
	//Write a program that finds all strings that include the letter i followed by one or more instances of the letter n.
	act3 := regexp.MustCompile(`in+`)
	fmt.Println(act3.FindString("input"))
	fmt.Println("Activity 4")
	//Write a program that finds all strings that include the letter i followed by one or two instances of the letter n.
	act4 := regexp.MustCompile(`in{1,2}`)
	fmt.Println(act4.FindString("input"))
	fmt.Println("Activity 5")
	//Write a program that finds all strings that include the letter i followed by three instances of the letter n.
	act5 := regexp.MustCompile(`in{3}`)
	fmt.Println(act5.FindString("innnput"))
	fmt.Println("Python Regular Expression")
	//Write a Python program to find sequences of lowercase letters joined with a underscore.
	act6 := regexp.MustCompile(`[a-z]+_[a-z]`)
	fmt.Println(act6.FindString("input_a"))
	//Write a Python program to find the sequences of one upper case letter followed by lower case letters
	act7 := regexp.MustCompile(`[A-Z]+[a-z]`)
	fmt.Println(act7.FindString("Input_a"))
	// Write a Python program that matches a string that has an 'a' followed by anything, ending in 'b'.
	//act8 := regexp.MustCompile(`[A-Z]+[a-z]`)
	//Write a Python program that matches a word at the beginning of a string
	//act9 := regexp.MustCompile(`[A-Z]+[a-z]`)
	//Write a Python program that matches a word at the end of string, with optional punctuation
	//act10 := regexp.MustCompile(`[A-Z]+[a-z]`)
	//Write a Python program that matches a word containing 'z'.
	//act11 := regexp.MustCompile(`[A-Z]+[a-z]`)
	//Write a Python program that matches a word containing 'z', not at the start or end of the word.
	//Write a Python program to match a string that contains only upper and lowercase letters, numbers, and underscores
	//Write a Python program where a string will start with a specific number
	//Write a Python program to remove leading zeros from an IP address
	//Write a Python program to check for a number at the end of a string.
	//18. Write a Python program to search the numbers (0-9) of length between 1 to 3 in a given string.
	//"Exercises number 1, 12, 13, and 345 are important"

	// 19. Write a Python program to search some literals strings in a string.
	// Sample text : 'The quick brown fox jumps over the lazy dog.'
	// Searched words : 'fox', 'dog', 'horse'

	// 20. Write a Python program to search a literals string in a string and also find the location within the original string where the pattern occurs.

	// Sample text : 'The quick brown fox jumps over the lazy dog.'
	// Searched words : 'fox'

	// 21. Write a Python program to find the substrings within a string.

	// Sample text :

	// 'Python exercises, PHP exercises, C# exercises'

	// Pattern :

	// 'exercises'

	// Note: There are two instances of exercises in the input string.

	// 22. Write a Python program to find the occurrence and position of the substrings within a string.

	// 23. Write a Python program to replace whitespaces with an underscore and vice versa.

	// 24. Write a Python program to extract year, month and date from an url.

	// 25. Write a Python program to convert a date of yyyy-mm-dd format to dd-mm-yyyy format.

	// 26. Write a Python program to match if two words from a list of words starting with letter 'P'.

	// 27. Write a Python program to separate and print the numbers of a given string.

	// 28. Write a Python program to find all words starting with 'a' or 'e' in a given string.

	// 29. Write a Python program to separate and print the numbers and their position of a given string.

	// 30. Write a Python program to abbreviate 'Road' as 'Rd.' in a given string.

	// 31. Write a Python program to replace all occurrences of space, comma, or dot with a colon.

	// 32. Write a Python program to replace maximum 2 occurrences of space, comma, or dot with a colon.

	// 33. Write a Python program to find all five characters long word in a string.

	// 34. Write a Python program to find all three, four, five characters long words in a string.

	// 35. Write a Python program to find all words which are at least 4 characters long in a string.

	// 36. Write a python program to convert camel case string to snake case string.

	// 37. Write a python program to convert snake case string to camel case string.

	// 38. Write a Python program to extract values between quotation marks of a string.

	// 39. Write a Python program to remove multiple spaces in a string. Go

	// 40. Write a Python program to remove all whitespaces from a string.
	// 41. Write a Python program to remove everything except alphanumeric characters from a string.

	// 42. Write a Python program to find urls in a string.

	// 43. Write a Python program to split a string at uppercase letters.

	// 44. Write a Python program to do a case-insensitive string replacement.

	// 45. Write a Python program to remove the ANSI escape sequences from a string.

	// 46. Write a Python program to find all adverbs and their positions in a given sentence.

	// Sample text : "Clearly, he has no excuse for such behavior."

	// 47. Write a Python program to split a string with multiple delimiters.

	// Note : A delimiter is a sequence of one or more characters used to specify the boundary between separate, independent regions in plain text or other data streams. An example of a delimiter is the comma character, which acts as a field delimiter in a sequence of comma-separated values.

	// 48. Write a Python program to check a decimal with a precision of 2.

	// 49. Write a Python program to remove words from a string of length between 1 and a given number.

	// 50. Write a Python program to remove the parenthesis area in a string.
	// Sample data : ["example (.com)", "w3resource", "github (.com)", "stackoverflow (.com)"]
	// Expected Output:
	// example
	// w3resource
	// github
	// stackoverflow

	// 51. Write a Python program to insert spaces between words starting with capital letters.

	// 52. Write a Python program that reads a given expression and evaluates it.
	// Terms and conditions:
	// The expression consists of numerical values, operators and parentheses, and the ends with '='.
	// The operators includes +, -, *, / where, represents, addition, subtraction, multiplication and division.
	// When two operators have the same precedence, they are applied to left to right.
	// You may assume that there is no division by zero.
	// All calculation is performed as integers, and after the decimal point should be truncated Length of the expression will not exceed 100.
	// -1 ? 10 9 = intermediate results of computation = 10 9

	// 53. Write a Python program to remove lowercase substrings from a given string.

	// 54. Write a Python program to concatenate the consecutive numbers in a given string.
	// Original string:
	// Enter at 1 20 Kearny Street. The security desk can direct you to floor 1 6. Please have your identification ready.
	// After concatenating the consecutive numbers in the said string:
	// Enter at 120 Kearny Street. The security desk can direct you to floor 16. Please have your identification ready.

	// 55. Write a Python program to convert a given string to snake case.
	// Sample Output:
	// java-script
	// gd-script
	// btw...-what-*-do*-you-call-that-naming-style?-snake-case?
	// 56. Write a Python program that takes any number of iterable objects or objects with a length property and returns the longest one.
	// Sample Output:
	// Orange
	// [1, 2, 3, 4, 5]
	// Java
	// Python
}

/*

A Quick Go Regex Cheat Sheet
Below is a list of commonly used regular expressions and Regex and their meaning. The list below is not comprehensive:

ab: a followed by b
a|b: a or b
a*: Zero or more a’s
a?: Zero or one a’s
a{2}: Two or more a’s
[ab], ^[ab]: Either a or b, except a or b (^ symbolises not, ie not a or b)
[a-z]: Any character a to z
[0-9]: Any number 0 to 9.
\d: Any digit. Similarly, a non-digit is \D or [^0-9]
\s: A whitespace character or [\t\n\f\r]. Similarly, \S is non-whitespace character or [^\t\n\f\r]
\w: A word character: [0-9A-Za-z_]. Similarly, \W means a non-word: [^0-9A-Za-z_]
[\t\n\f\r\v]: Means a tab=\011, newline=\012, form feed=\014, carriage return=\015, vertical tab=\013 respectively.
\123: Octal character upto exactly three digits
\x9E: Exactly two digit hex character
\A or ^: Beginning of the text
$ or \z: End of the text
i: Case insensitive

Note: To match special characters, it must be escaped with a backslash character. For example, to match a $, prefix it with a backslash – \$.





import (
	"fmt"
	"regexp"
)

func main() {
/*
str := "Golang expressions example"

regexp,_ := regexp.Compile("Gola([a-z]+)g")

fmt.Println(regexp.FindString(str))



str := "Golang expressions example"

    regexp,_ := regexp.Compile("Gola([a-z]+)g")

    fmt.Println(regexp.FindString(str))



/*

	match, err := regexp.MatchString(`[^0-9]`, "1234567")
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(match)

      text := "This is *$ // sample ## %% text"

	reg := regexp.MustCompile(`[^a-zA-Z0-9]`)
	fmt.Println("true|false", reg.MatchString(text))

	strs := reg.FindAllString(text, -1)

	for _, e := range strs {
		fmt.Println(e)
	}

*/
