package july2022

import (
	"fmt"
)

func sum(n []int, o int) int {
	arr := n[o:]
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
func halves(n int) {
	half := n % 2
	if half == 0 {
		fmt.Println("1, true")
	} else {
		fmt.Println("0, false")
	}
}
func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
func greatest(arr []int) int {
	great := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > great {
			great = arr[i]
		}
	}
	return great
}
func makeOddGenerator() func() int {
	i := 1
	return func() int {
		i += 2
		return i
	}
}

//borrowed
// f =
func recursiveFib(f, s, n int) {
	var temp int = 0
	if n > 0 {
		fmt.Println(f)
		temp = f + s
		f = s
		s = temp
		recursiveFib(f, s, n-1)
	}
}

//cache fib from class
func fib_memo(n int) int {
	cache := make(map[int]int)
	return _fib(n, cache)
}

//don't call _ functions directly
func _fib(n int, cache map[int]int) int {
	if n < 2 {
		return n
	}
	if i, ok := cache[n]; ok {
		return i
	}
	cache[n] = _fib(n-1, cache) + _fib(n-2, cache)
	return cache[n]
}

// func fib_it(n int) int {
// 	a := 1
// 	b := 2
// 	for n > 2 {
// 		a, b = b, a+b
// 	}
// }
func pointer(n []int) {
	var numbersptr [3]*int
	for i := 0; i < len(numbersptr); i++ {
		numbersptr[i] = &n[i]
	}
}

func Day5() {
	// 1)sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go? Function Signature what is entered and the return type.
	// 2)Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).
	// 3). Write a function with one variadic parameter that finds the greatest number in a list of numbers.
	// 4)Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.
	// 5)The Fibonacci sequence is defined as: 0, fib(0) = fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).
	q1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pos := 5
	fmt.Println("The sum of the array slice is: ", sum(q1, pos))
	halves(2)
	halves(3)
	fmt.Println("The fibonacci number of 5 is: ", fibonacci(5))
	q3 := []int{5, 7, 15, 2, 684, 8, 2}
	fmt.Println("The greatest number of the array", q3, "is: ", greatest(q3))
	f := makeOddGenerator()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	//recursiveFib(f, s, n-1)
	fmt.Println(fib_memo(10))
	//Array of Pointers
	numbers := []int{100, 1000, 10000}
	// var i int
	// for i = 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	pointer(numbers)
	// 	fmt.Println(numbersptr)     // print the pointer array
	// 	fmt.Println(*numbersptr[0]) // print the values the pointers point to
	// 	fmt.Println(*numbersptr[1])
	// 	fmt.Println(*numbersptr[2])
	// 	*numbersptr[0] = 200
	// 	fmt.Println(numbers)
}

// func main() {

// f := evenGenerator()
// fmt.Println(f())
// fmt.Println(f())
// fmt.Println(f())
// fmt.Println(f())
// }

// func evenGenerator() func() int {
// i := 0
// return func() int {
// 	i += 2
// 	return i
// }
// }

//Nodes
// package main
// import "fmt"
// type Node struct{
//     Value int
// }

// type Stack struct{
//     nodes[] *Node
//     count int
// }
// func NewStack() *Stack{
//     return &Stack{}
// }
//   func (s *Stack)  Push(n *Node){
//          s.nodes=append(s.nodes[:s.count],n)

//          s.count++//pont to next position
//     }
// func (s *Stack)  Pop()  *Node{
//          if s.count==0{ return nil}
//          s.count--
//         return  s.nodes[s.count]
//     }
// func (n  *Node) String()string{
//     return fmt.Sprint(n.Value)
// }
// func main(){
//    s:=NewStack()
//    s.Push( &Node{10} )
//    s.Push( &Node{12} )
//    s.Push( &Node{14} )
//    s.Push( &Node{16} )

//     fmt.Println(s.Pop()  ,s.Pop(),s.Pop(),s.Pop())

// }
