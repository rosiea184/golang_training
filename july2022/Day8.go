package july2022

import (
	"fmt"
	"math"
)

func changebyval(x int) {
	x = x + 10
	fmt.Println(x)
}
func changebyAddress(x *int) {
	*x = *x + 10
	fmt.Println(*x)
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

func A() {
	fmt.Println("Calling A")
}
func B(t int) {
	fmt.Println("Calling B", t)
}
func C(t int) {
	fmt.Println("Calling C", t)
}
func D(A string, f float32) {
	fmt.Println("Calling D", A, " ", f)
}

func Day8() {
	f := []interface{}{A, B, D}
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

	op := math.Abs
	a := []float64{1, -2}
	b := Map(op, a)
	fmt.Println(b) //[1 2]
	// pointer8()
	// a := 10
	// changebyval(a)
	// changebyAddress(&a)

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
