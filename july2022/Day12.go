package july2022

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
1)Create a slice of 20 of type int and take 20 number create 4 goroutines to take average of every 5 elements ex[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20] ,so 1 goroutine should give average of first 5 elements and next for another 5 so on
*/
func adv1(a []int, c chan int) {
	sum := 0
	for _, d := range a {
		sum += d
	}
	c <- sum
}
func adv2(a []int, c chan int) {
	sum := 0
	for _, d := range a {
		sum += d
	}
	c <- sum
}
func adv3(a []int, c chan int) {
	sum := 0
	for _, d := range a {
		sum += d
	}
	c <- sum
}
func adv4(a []int, c chan int) {
	sum := 0
	for _, d := range a {
		sum += d
	}
	c <- sum
}

/*
2)Take a Paragraphs of text Max word  up to 200 words in a Slice or String and then Pass each  Sentence(up to full stop) to another go routine/routines   1)print it in reverse order(1 goroutine) 2 )count its words(2 goroutine)
*/
func printReverse(b []string, results chan []string) {
	var reverseString []string
	for i := 0; i < len(b); i++ {
		wordString := strings.Split(b[i], " ")
		for j := len(wordString) - 1; j < 0; j-- {
			fmt.Println(wordString[j])
			//time.Sleep(1 * time.Millisecond)
			reverseString = append(reverseString, wordString[j])
		}
	}
	results <- reverseString

}
func countWords(b string, e chan string) {
	wordString := strings.Split(b, " ")
	length := strconv.Itoa(len(wordString))
	e <- length
	time.Sleep(1 * time.Millisecond)
}

func Day12() {
	//Activity 1
	// rand.Seed(time.Now().UnixNano())
	// var a []int
	// for i := 0; i < 20; i++ {
	// 	a = append(a, rand.Intn(9)+1)
	// }
	// c := make(chan int)
	// go adv1(a[:5], c)
	// go adv2(a[6:10], c)
	// go adv3(a[11:15], c)
	// go adv4(a[16:20], c)
	// w, x, y, z := <-c, <-c, <-c, <-c
	// fmt.Println("1st set average: ", w)
	// fmt.Println("2nd set average: ", x)
	// fmt.Println("3rd set average: ", y)
	// fmt.Println("4th set average: ", z)
	//-----------------------------------
	//Activity 2
	// paragraph := "This is a sentance. This is another sentance."
	// //fmt.Println(paragraph)
	// splitSentence := strings.Split(paragraph, ".")
	// sentanceL := len(splitSentence)
	// //e := make(chan int, sentanceL)
	// results := make(chan []string)
	// //it keeps deadlocking
	// // for i := 0; i < sentanceL; i++ {
	// // 	sentance := splitSentence[i]
	// go printReverse(splitSentence, results)
	// //time.Sleep(1 * time.Second)
	// //go countWords(sentance, e)
	// //e <- sentanceL
	// s := <-results //, <-e
	// for a := 1; a <= sentanceL; a++ {
	// 	<-results
	// }
	// fmt.Println(s)
	// //time.Sleep(1 * time.Millisecond)
	// //fmt.Println(t)
	// //}
	// //close(e)
	// time.Sleep(5 * time.Second)
	//Activity 3
	e := make(chan int)
}

/*
3)Create a goroutine/channels to send and receive structure data type  free to design for any purpose                                                                                                                                                                              4)Create Two go routine where routine 1 generate random number and append them in  slice where another goroutine sort the slice at the same time . print the slice after every append and sorted array at the same time  side by side
*/
