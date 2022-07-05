package main

import (
	"fmt"
	"math/rand"
	"time"
	//"strconv"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Float32()
	fmt.Println(x)
	fmt.Println(rand.Float64())
	fmt.Print(rand.Intn(5) + 5)

}
