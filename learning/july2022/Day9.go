package july2022

import (
	"fmt"
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

func Day9() {
	fmt.Println("Day 9 Notes")
}
