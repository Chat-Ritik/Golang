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

// Model for Couse-file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	Author      *Author `json:"author"`
	CoursePrice int     `json:"price"`
}

// Model for Author details
type Author struct {
	FullName string `json:"name"`
	Website  string `json:"website"`
}

// Fake database - slice
var courses []Course

// middleware or helper-file --- to check courseId which is a part of struct, so define a method
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

// controllers-file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to our page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Alll Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")
	//grab an id of the course demanded
	params := mux.Vars(r)
	//loop through courses,find matching id and return the Response
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No corse found with given ID")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")
	//if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	//if empty {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	//generate unique ID,string
	//append course into Courses
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
	//check if coursename is matching, ask to change
	for _, c := range courses {
		if course.CourseName == c.CourseName {
			json.NewEncoder(w).Encode("Please change the course name as it already exists.")
			return
		}
	}
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")
	//first grab id from Request
	params := mux.Vars(r)
	//loop, id remove,add with new ID
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No such CourseID found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")
	//grab id to Delete
	params := mux.Vars(r)
	//loop,id,remove(index,index+1)
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("The course is deleted.")
			break
		}
	}
}

func main() {
	fmt.Println("API-Build.in")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{CourseID: "2", CourseName: "ReactJS", Author: &Author{FullName: "Simon Hughes", Website: "hughesops.dev"}, CoursePrice: 399})
	courses = append(courses, Course{CourseID: "3", CourseName: "NodeJS", Author: &Author{FullName: "John David", Website: "johndavidcode.dev"}, CoursePrice: 599})
	courses = append(courses, Course{CourseID: "6", CourseName: "Mern stack", Author: &Author{FullName: "Jacob Wins", Website: "jacobwin.dev"}, CoursePrice: 499})
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}
