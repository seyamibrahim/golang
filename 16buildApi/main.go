package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course -file
type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	// Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

var courses = []Course{{"adfslk", "React", 199}, {"flaskdjf", "Go", 299}}

// middleware , helper -file

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

	r := mux.NewRouter()
	// routes
	r.HandleFunc("/", ServeHome).Methods("GET")

	r.HandleFunc("/getallcourses", GetAllCourses).Methods("GET")
	r.HandleFunc("/getallcourses/{id}", getOneCourse).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Api Learning</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)
	id := params["id"]
	// loop through courses, find matching id and return the response
	// idd := params["id"]

	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id ")

}
