package july2022

import (
	"fmt"
	"math"
	//"sort"
)

func changebyval(x int) {
	x = x + 10
	fmt.Println(x)
}
func changebyAddress(x *int) {
	*x = *x + 10
	fmt.Println(*x)
}
func pointer8() {
	a := 43

	fmt.Println(a)  //43
	fmt.Println(&a) //

	b := &a

	fmt.Println(b)  //0x20818a220
	fmt.Println(*b) //43

	*b = 42 //b says
}

func map8() {
	var myGreeting map[string]string

	myGreeting["A"] = "ABC" //Key = A values "ABC"
	//fmt.Println(myGreeting == nil)
	//panic: assignment to entry in nil map
	myGreeting["Tim"] = "Good Morning."
	myGreeting["Jenny"] = "Bonjour."
	fmt.Println(myGreeting)
	for i, j := range myGreeting {
		fmt.Println(i, j)
	}
	//to get keys out in alphabetical order you have to sort the keys.
	//Printing does it automatically. Going through the loop will not.
}

// type Person struct {
// 	First string
// 	Last  string
// 	Age   int
// }

// func (p Person) print() {
// 	fmt.Println(p)
// 	fmt.Printf("%+v \n", p)
// }

type Operator func(x float64) float64

func Map(op Operator, a []float64) []float64 {
	res := make([]float64, len(a))
	for i, x := range a {
		res[i] = op(x)
	}
	return res
}
func operatorExample() {
	op := math.Abs
	a := []float64{1, -2}
	b := Map(op, a)
	fmt.Println(b) //[1 2]
}

func A() {
	fmt.Println("Calling A")
}
func B(t int) {
	fmt.Println("Calling B", t)
}
func C(A string, f float32) {
	fmt.Println("Calling C", A, " ", f)
}
func interfaceExample() {
	//interface
	f := []interface{}{A, B, C}
	f[0].(func())()
	f[1].(func(int))(15)
	f[2].(func(string, float32))("Hello ", 15.15)
	// var slicefn []func()
	// slicefn = append(slicefn, A)
	// slicefn = append(slicefn, B)
	// slicefn = append(slicefn, C)
	// for _, f := range slicefn {
	// 	f()
	// }
}

func mapfromfn(item map[string]float64) map[string]float64 {
	db := make(map[string]float64)
	for k, v := range item {
		db[k] = v * 2
	}
	return db
}
func mapfromfnMain() {
	//mapping
	p := make(map[string]float64)
	p["Iphone"] = 99256.50
	p["Android"] = 5000
	p["Xamarin"] = 8000
	updatadata := mapfromfn(p)
	for key, val := range updatadata {
		fmt.Printf("%s %v \n", key, val)
	}
} /*
type shape interface {
	area() float64
}
type square struct{
	side float64
}
type cube struct{
	side float64
}
func (z square) area()float64{
	return z.side * z.side
}
func info(z shape){//Interface as parameter
	fmt.Println(z)
	fmt.Println(z.area())
}
func interfaceE2(){
	s:=square{10}
	c:= circle{5}
	info(s)
	info(&c)
}*/

func Day8() {

	//sort
	// n:= []int{7,4,8,2,9,19,12,32,3}
	// fmt.Println(n)
	// sort.Sort(sort.Reverse(sort.IntSlice(n)))
	// fmt.Println(n)
	/*
			   2) Take 20 (any count) from console between 1 -100 ) then print the
		2) Take 20 (any count) from console between 1 -100 ) then print the
			   2) Take 20 (any count) from console between 1 -100 ) then print the
			   summary like no between 1- 10 Count-5     11 -20 count -7    21-30
		summary like no between 1- 10 Count-5     11 -20 count -7    21-30
			   summary like no between 1- 10 Count-5     11 -20 count -7    21-30
			   count -10  etc
			   3) Create a map ex map[string]int{"orange": 5, "apple": 7,	"mango": 3,
			    "strawberry": 9} ,sort the map based on key length in asecending order
	*/

	//maps activity keyword search
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

	//map8()

	pointer8()
	// a := 10
	// changebyval(a)
	// changebyAddress(&a)

}
