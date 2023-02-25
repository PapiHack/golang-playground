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
	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON()  {
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

func DecodeJSON()  {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN Bootcamp",
		"Price": 199,
		"website": "learncodeonline.in",
		"tags": ["full-stack", "js"]
	}
	`)
	var lcoCourse Course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid!")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid !")
	}

	// Some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v an Type is : %T\n", k, v, v)
	}
}
