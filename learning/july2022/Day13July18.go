package july2022

import (
	//"fmt"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var port = flag.String("port", "8000", "the port number")
var tz = flag.String("tz", "8000", "time zone")

// call like ./clock1 -port 8010 -tz=Asia/Tokyo
func Day13() {
	flag.Parse()
	os.Setenv("TZ", *tz)
	listener, err := net.Listen("tcp", "localhost:"+*port)
	fmt.Printf("Listening on port %s\n", *port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		fmt.Println("Received connection from: ", c.RemoteAddr())
		time.Sleep(1 * time.Second)
	}
}

//hint for wall clocks
/*Ipackage main

import (
	"io"
	"log"
	"net"
	"os"
)

// ./wallofclocks NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
func main() {
	// os.Args holds the command-line arguments, starting with the program name.
	for _, place := range os.Args[1:] {
		// place is gonna look like NewYork=localhost:8010
		// get the localhost:port part
		// make a connection to that localhost:port
	}
}

func timeAt(address string) {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	shouldCopy(os.Stdout, conn)
}

func shouldCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
*/
//net notes
/*
func Day13() {
	conn, err := net.Dial("tcp", "localhost:8000")
	// listener, err := net.Listen("tcp", "localhost:8000")
	// //http 80
	// //https 443
	// //hss 22
	if err != nil {
		log.Fatal(err)
	}
	shouldCopy(os.Stdout, conn)
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Print(err) // connection aborted
	// 		continue
	// 	}
	// 	handleConn(conn) // handle one connection at a time
	// }
}
func shouldCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
func handleConn(c net.Conn) {
	defer c.Close() //do this at end of function
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return //client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

/* Spinner time function
func Day13() {
	const n = 45
	go spinner(10 * time.Second)
	fibN := fib(n) // naive and slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	//spinner runs until main function is done
	spin := [4]string{"|", "/", "-", "\\"}
	j := true
	i := 0
	for j {
		fmt.Printf("\r%s", spin[i%4])
		time.Sleep(250 * time.Millisecond)
		i++
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

/*
//i := time.Now()
//time.Since(i) < delay
//fmt.Printf("3\r") clears back to the beginning
*/
