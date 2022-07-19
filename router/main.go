package main

import (
	//"mymodule/router"
	//"golang-rest-api/router"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/gorilla/handlers"
	//"github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handler 1")
}
func main() {
	//----------------------------------
	//Router Stuff
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/us", handler1).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
