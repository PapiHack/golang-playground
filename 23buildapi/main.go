package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	return course.Id == "" && course.Name == ""
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
