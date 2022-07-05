package july2022

import (
	"fmt"
	"math/rand"
	"time"
	//"strconv"
)

func Day2() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Float32()
	fmt.Println(x)
	fmt.Println(rand.Float64())
	fmt.Print(rand.Intn(5) + 5)

}
