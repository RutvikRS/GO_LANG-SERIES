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

//Model for courses - file

type Course struct {
	CourseID    string  `json:"courseID"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fakeDB
var courses []Course

//middleware, helper - file

func (c *Course) isEmpty() bool {
	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - Course Controller ")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseID: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Rutvik Sakpal", Website: "lco.dev"}})
	courses = append(courses, Course{CourseID: "4", CourseName: "NodeJS", CoursePrice: 499, Author: &Author{Fullname: "Yogesh More", Website: "lco.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getonecourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearncodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses,find matching id and return the response

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about - {}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		fmt.Println("err", err)
	}
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate unique id, string
	//append course into courses

	s := rand.NewSource(time.Now().UnixNano())
	h := rand.New(s)
	course.CourseID = strconv.Itoa(h.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req

	params := mux.Vars(r)

	// loop, id, remove, add with my ID

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

	//TODO: send a response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop,id,remove(index,index+1)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}
