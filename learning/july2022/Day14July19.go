package july2022

import (
	//"mymodule/router"
	//"golang-rest-api/router"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	//"github.com/gorilla/handlers"
	//"github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func handler1(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Handler 1")
}
func Now(w http.ResponseWriter, _ *http.Request) {
	//need to get different time zone
	now := time.Now()
	p := make(map[string]string)
	p["now"] = now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

/*
1) respond client to a date and time ,not for server but specifically client timezone as restAPI  GET Request

Assignment 2 )   Write  a go program to get the client IP
and reflect/echo back to client page ,after this add
feature to server so that server could block to a
particular client with a message add feature to store all
past access like date and time to access to the server and
display a message that  clientIP:xxxxxxxxxxxx  your last
access to server is  Date  :  xxxx/xx/xxxx time :xxxxxxx
add feature to store all past access like date and time to
access to the server and display a message that
clientIP:xxxxxxxxxxxx  your last access to server is
Date  :  xxxx/xx/xxxx time :xxxxxxx


Hint  : func ReadUserIP(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-Ip")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarded-For")
    }
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
    }
    return IPAddress
}
X-Real-Ip - fetches first true IP (if the requests sits
	behind multiple NAT sources/load balancer)

X-Forwarded-For - if for some reason X-Real-Ip is blank
and does not return response, get from X-Forwarded-For

Remote Address - last resort (usually won't be reliable as
this might be the last ip or if it is a naked http request
to server ie no load balancer)

assignment 3) Convert json data article := `{"id": "BM-1347", "title": "The underage storm", "Content": "The creatives' careers can easily get uncreative but yet creative...", "Summary": "Seeking freedom"}`               ...........into  map and send it to client as GET http request

*/
func Day14() {
	//----------------------------------
	//Router Stuff
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/us", handler1).Methods("GET")
	r.HandleFunc("/now", Now).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
