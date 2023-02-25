package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()
}

func EncodeJson()  {
	lcoCourses := []Course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"},},
		{"MERN Bootcamp", 199, "learncodeonline.in", "bcd123", []string{"full-stack", "js"},},
		{"Angular Bootcamp", 299, "learncodeonline.in", "hit123", nil,},
	}

	// Package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
