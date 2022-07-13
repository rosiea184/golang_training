package july2022

import (
	"fmt"
	"time"
	//"sync"
)

//July 8 & 13

// func foo() {
// 	for i := 0; i < 45; i++ {
// 		fmt.Println("Foo:", i)
// 		//time.Sleep(1 * time.Millisecond)
// 	}
// 	wg.Done()
// }
// func bar() {
// 	for i := 0; i < 45; i++ {
// 		fmt.Println("bar:", i)
// 		//time.Sleep(1 * time.Millisecond)
// 	}
// 	wg.Done()

// }

//var wg sync.WaitGroup

//allows a go to complete it's cycle of execution

func Day10() {

	//Concurency. time is used so things run at the same time.
	//Good for multitasking
	// wg.Add(2)
	// go foo()
	// go bar()
	// wg.Wait()

	//time.Sleep(1 * time.Second) //allows other go to executes
	//ch := make(chan string)
	fmt.Println("Main")
	// go Sendme(ch)
	// fmt.Println(<-ch) //Reading ch
	//fmt.Println("Main End")
	// c := make(chan int)
	// go func() {
	// 	for i := 0; i <= 100; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()
	// for n := range c {
	// 	fmt.Println(n)
	// }
	//Practice Concurrency
	go countUp(100)
	go countDown(100)
	time.Sleep(5 * time.Second)
	fmt.Println("Main ended")
}
func countUp(limit int) {
	for i := 0; i <= limit; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("Counting up:", i)
	}
}
func countDown(limit int) {
	for i := limit; i >= 0; i-- {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("Counting down:", i)
	}
}

//ch chan <- string ch is only used for writing || ch <- chan string reading
// func Sendme(ch chan<- string) {
// 	fmt.Println("SendMe Wait")
// 	time.Sleep(time.Second * 5)
// 	ch <- "Send me done"
// }
