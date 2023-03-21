package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const PORT = 4000

// Model for course - file

type Course struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Author *Author `json:"author"`
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
	fmt.Println("API - LearnCodeOnline.in")
	fmt.Printf("Server listening on port %v...\n", PORT)

	router := mux.NewRouter()

	// Seeding our fake DB
	courses = append(courses, Course{
		Id: "2", 
		Name: "ReactJS",
		Price: 299,
		Author: &Author{
			Fullname: "Papi MBAYE",
			Website: "lco.dev",
		},
	})
	courses = append(courses, Course{
		Id: "4", 
		Name: "MERN",
		Price: 199,
		Author: &Author{
			Fullname: "Papi MBAYE",
			Website: "go.dev",
		},
	})

	// Handle routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/courses", createOneCourse).Methods("POST")
	router.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")
	router.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/courses", deleteAllCourse).Methods("DELETE")

	// Listening to a port
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(PORT), router))
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
			return
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
	
	var newCourse Course
	_ = json.NewDecoder(request.Body).Decode(&newCourse)
	if newCourse.IsEmpty() {
		json.NewEncoder(writer).Encode("No data inside the JSON.")
		return
	}

	// TODO: check if course name already exist

	for _, course := range courses {
		if course.Name == newCourse.Name {
			json.NewEncoder(writer).Encode("Course name '" + newCourse.Name + "' already exist.")
			return
		}
	}
	
	rand.NewSource(int64(rand.Intn(100)))
	newCourse.Id = strconv.Itoa(rand.Intn(100))
	courses = append(courses, newCourse)
	json.NewEncoder(writer).Encode(newCourse)
}

func updateOneCourse(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("UPDATE ONE COURSE")
	writer.Header().Set("Content-Type", "application/json")

	if request.Body == nil {
		json.NewEncoder(writer).Encode("Please send some data.")
		return
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
			return
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
			return
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
