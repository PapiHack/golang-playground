package main

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