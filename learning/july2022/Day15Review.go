package july2022

import (
	"fmt"
	"time"
)

func countUp1(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}
func countDown1(ch chan int) {
	for i := 100; i > 0; i-- {
		ch <- i
	}
	close(ch)
}

func Day15() {
	//Concurancy Channels
	//1-100
	ch1 := make(chan int)
	ch2 := make(chan int)
	go countUp1(ch1)
	go countDown1(ch2)
	b := <-ch1
	d := <-ch2
	for i := 0; i < 100; i++ {
		fmt.Println(b)
		fmt.Println(d)
		b = <-ch1
		d = <-ch2
	}
	time.Sleep(4 * time.Second)
}

/*
var (
	group   sync.WaitGroup
	channel chan string
)

func Forword() {
	for i := 0; i <= 10; i++ {
		channel <- fmt.Sprintf("Forword i=%v", i) // send message
		time.Sleep(time.Millisecond)
	}

	fmt.Println("Forword finish its task ")
	group.Done()
}
//Printng both the values
func Reverse () {
	for i := 10; i >= 0; i-- {
		if s,ok := <- channel; ok { // receive message
			fmt.Println(s) //Print the value of Forword

			fmt.Printf("   Reverse i=%v\n", i)// Print the value of Reverse
			time.Sleep(time.Millisecond)
		}
	}

	fmt.Println(" Reverse finished task")
	group.Done()
}

func main() {
	channel = make(chan string)
	group.Add(2)
	go Forword()
	go Reverse()
	group.Wait()
	close(channel) // close it so nothing is waiting
	for s := range channel {
		fmt.Println(s)
	}
}
*/
