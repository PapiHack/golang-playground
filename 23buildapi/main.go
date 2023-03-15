package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course - file

type Course struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Author *Author
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// Fake DB
var courses []Course

// Middleware, helper - file
func (course *Course) IsEmpty() bool {
	// return course.Id == "" && course.Name == ""
	return course.Name == ""
}

func main() {
	
}

// Controllers - file

// Serve home route
func serveHome(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>Welcome to API powered by LearnCodeOnline!</h1>"))
}

func getAllCourses(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("GET ALL COURSES")
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(courses)
}

func getOneCourse(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("GET ONE COURSE")
	writer.Header().Set("Content-Type", "application/json")

	// Grab id from request
	params := mux.Vars(request)

	// Loop through courses, find matching id and return resoponse
	for _, course := range courses {
		if course.Id == params["id"] {
			json.NewEncoder(writer).Encode(course)
		}
	}
	json.NewEncoder(writer).Encode("No Course found with id.")
}

func createOneCourse(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("CREATE ONE COURSE")
	writer.Header().Set("Content-Type", "application/json")

	if request.Body == nil {
		json.NewEncoder(writer).Encode("Please send some data.")
	}
	
	var course Course
	_ = json.NewDecoder(request.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(writer).Encode("No data inside the JSON.")
	}
	
	rand.NewSource(int64(rand.Intn(100)))
	course.Id = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(writer).Encode(course)
}

func updateOneCourse(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("UPDATE ONE COURSE")
	writer.Header().Set("Content-Type", "application/json")

	if request.Body == nil {
		json.NewEncoder(writer).Encode("Please send some data.")
	}

	params := mux.Vars(request)
	for index, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse Course
			json.NewDecoder(request.Body).Decode(&updatedCourse)
			updatedCourse.Id = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(writer).Encode(updatedCourse)
		}
	}

	json.NewEncoder(writer).Encode("No course found with id.")
	
}

func deleteOneCourse(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("DELETE ONE COURSE")
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	for index, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(writer).Encode("Course with id #" + course.Id + " was deleted.")
			break
		}
	}

	json.NewEncoder(writer).Encode("No course found with id.")
}

func deleteAllCourse(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("DELETE ALL COURSES")
	writer.Header().Set("Content-Type", "application/json")
	courses = []Course{}

	json.NewEncoder(writer).Encode("All course was deleted.")
}
