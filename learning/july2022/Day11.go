package july2022

import (
	"fmt"
)

//July 14th 2022
//More concurency, deadlock
//Always close the channel.
func Day11() {
	//call num gen
	//see if num is prime
	c := make(chan int)
	go func() {
		for i := 0; i <= 9; i++ {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
}

/*import (
	"fmt"
	"sync"
    "math/rand"
    "sync/atomic"
	"time"
)
var counter int64
var wg sync.WaitGroup
func incrementor(s string ){
    for i:=0;i<=10 ;i++{
              time.Sleep(time.Duration(rand.Intn(3))* time.Millisecond)
             // counter++
              atomic.AddInt64(&counter,1)
              fmt.Println(s , i , "Countert " , atomic.LoadInt64(&counter))

    }
    wg.Done()
}
func main(){
    wg.Add(2)
    go incrementor("Routine 1")
    go incrementor("Routine 2")
    wg.Wait()
    fmt.Println( "Countert " , counter )
}
*/

/* Prime Checker
func inc(s string ) chan int {
    out:=make (chan int)
       go  func (){

              for i:=0 ; i<=5 ;i++{
                  out<-i
                  fmt.Println(i)
              }
              close(out)
        } ()

    return out
}
func puller(c chan int  )chan int{
    out:=make(chan int)
    go func (){
           var sum int
           for n:=range c{
               sum+=n
             }
             out<-sum
             close(out)
    } ()
    return out
}

func main(){
    n1:=inc("Genarate1")
    p1:=puller(n1)

    v,ok :=<-p1
    fmt.Println("Sum " , v, ok  ,cap(p1))

}*/

/*
func Mix(s, r chan struct{}) {
	for {
		select {
		case d1 := <-s:
			fmt.Println("Recieved from Bob", d1)
			r <- struct{}{}
		case d2 := <-r:
			fmt.Println("Recieved from Johan", d2)
			s <- struct{}{}
		}
	}
}
func Day11() {
	s := make(chan struct{}, 1)
	r := make(chan struct{}, 1)
	s <- struct{}{}
	go Mix(s, r)
	time.Sleep(1 * time.Second)
}
*/
/*
func Day11() {
	send := make(chan int, 1)
	reciev := make(chan int, 1)
	go sender(send, reciev)
	go reciever(reciev, send)
	send <- 1
	select {}
}
func sender(snd <-chan int, recv chan<- int) { //(snd <-chan int(Reading), recv chan <- int (writing)
	inform := <-snd
	fmt.Println("Get information from Send ", inform)
	time.Sleep(1 * time.Second)
	recv <- inform + 1
}
func reciever(recv <-chan int, snd chan<- int) {
	for {
		inform := <-recv //information is recieved
		fmt.Println("Information from Reciever ", inform)
		time.Sleep(1 * time.Second)
		snd <- inform + 1
	}
}
*/
/*
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// Creating a channel without a buffer.
	// This will make the calling go routine to wait until a recevier is up and processing the data.
	// If we wanted a buffer we could specify it this way 'ch1 := make(chan int,5)' which will make a channel with a buffer of 5.
	// With a buffer of 5 the calling go routine will be able to put 5 int's in the channel,
	// before it have to wait for the receiver to process the data on the channel
	ch1 := make(chan int)
	// no wait group adding here, since the 'get' function will not be finnished before 'put' is done, and the channel is closed either way.
	go put(ch1)// writing
	// Adding 1 to the waitgroup since there is only 1 go routine we will wait on
	wg.Add(1)
	go get(ch1)
	// Will wait here until the go routine are finnished and the waitgroup is zero before it continues
	wg.Wait()
}
func put(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
	}
	// Closing the channel on the sender when the job is done.
	// NB : The sender shall allways close the channel, never the receiver.
	close(c)
}
func get(c chan int) {
	for {
		// Using two variables with channels, the 2nd variable will have the state of the channel as type bool
		myVar, ok := <-c // reading    <-ch    writing  ch<- 12
		fmt.Printf("The value of 'ok' = %v, and the type is %T\n", ok, ok)
		if ok {
			fmt.Println("content of myVar = ", myVar, "and the content of c = ", c)
			time.Sleep(time.Millisecond * 500)
		} else {
			// The channel is closed since 'ok=false', and we do the logic to decrement the waitgroup, and leave the function
			fmt.Println("Channel was closed, returning for loop")
			// Decrement the waitgroup value by 1 by calling wg.Done()
			wg.Done()
			return
		}
	}
}
*/
