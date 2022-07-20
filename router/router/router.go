package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	//"fmt"
	"github.com/gorilla/mux"
	//"github.com/gorilla/handlers"
	//"github.com/sirupsen/logrus"
)

type task struct {
	ID      int `json:ID`
	Name    string
	Content string
}
type allTasks []task //list of tasks

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some content",
	},
}

type allEmployees []employee

var employees = allEmployees{}

type employee struct {
	ID       int     `json:ID`
	Name     string  `json:Name`
	BasePay  float64 `json:Base Pay`
	Dept     string  `json:Department`
	HRA      float64 `json:HRA`
	GrossPay float64 `json:Gross Pay`
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "To Add Employee please enter name, base pay, and department /n")
	json.NewEncoder(w).Encode(employees)
}
func createEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee employee
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &newEmployee)

	newEmployee.ID = len(employees) + 1
	if newEmployee.Dept == "IT" {
		newEmployee.HRA = newEmployee.BasePay * .1
		newEmployee.GrossPay = newEmployee.BasePay + newEmployee.HRA
	} else if newEmployee.Dept == "Security" {
		newEmployee.HRA = newEmployee.BasePay * .14
		newEmployee.GrossPay = newEmployee.BasePay + newEmployee.HRA
	}
	employees = append(employees, newEmployee)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEmployee)

}

/*
Request by client (POST HTTP )
id
Name
Basic Pay   50000
Dept_Name   IT/Security
HRA         IT =10%BP  Security 14%BP
Gross Pay       0
------------------
Response by server
id
Name
Basic Pay   50000
HRA         10000
Gross Pay   60000
*/

//HTTP GET
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Fprintf(w, "Task with ID %v has been removed successfully", taskID)
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updateTask task

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &updateTask)

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			updateTask.ID = taskID
			tasks = append(tasks, updateTask)

			fmt.Fprintf(w, "The task with id %v has been updated successfully", taskID)
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

//ConfigureRouter setup the router
func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/ping", healthCheck).Methods("GET")
	router.HandleFunc("/employee", getEmployee).Methods("GET")
	router.HandleFunc("/employee", createEmployee).Methods("POST")
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}

// func Router() {
// 	router := router.ConfigureRouter()
// 	fmt.Println(router)
// }
