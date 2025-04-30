package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course -file
type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"-"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

var courses = []Course{{"1", "React", 199, &Author{"seyam", "google.com"}}, {"2", "Go", 299,&Author{"seyam", "go.dev"}}}

// middleware , helper -file

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

	r := mux.NewRouter()
	// routes
	r.HandleFunc("/", ServeHome).Methods("GET")

	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deletOneCourse).Methods("DELETE")

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
	for k, val := range params {
		fmt.Printf("%v :  %v", k, val)
	}
	id := params["id"]
	// loop through courses, find matching id and return the response
	// idd := params["id"]

	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{
		"error":     "course is not found",
		
	})

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Empty data")
		return
	}
	// generate unique id , string

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter,r *http.Request){

	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id := params["id"]

	// loop ,id, remove , add with my Id
	for i , course := range courses{
		if(course.CourseId == id){

			// remove
			courses = append(courses[:i], courses[i+1:]...)
			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = id
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}
		
	}

	json.NewEncoder(w).Encode("Course Not Found")
}

func deletOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Delete one Course")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	for i , course := range courses{
		if(course.CourseId == id){
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"deleted course_id" : id,
				"msg" : "Course Deleted",

			})
			return
		}
	}
	json.NewEncoder(w).Encode("Course Not found")
}