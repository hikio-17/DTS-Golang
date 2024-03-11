package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `
	[
		{
			"full_name": "Fajri Muhammad Tio",
			"email": "airel@gmail.com",
			"age": 23
		},
		{
			"full_name": "Hikma Nelda",
			"email": "airel@gmail.com",
			"age": 23
		}
	]
	`

	// 1.
	// var result Employee

	// 2.
	// var result map[string]interface{}

	// 3.
	// var temp interface{}
	// var err = json.Unmarshal([]byte(jsonString), &temp)
	// var result = temp.(map[string]interface{})

	var results []Employee

	var err = json.Unmarshal([]byte(jsonString), &results)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range results {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}

	// fmt.Println("full_name:", result["full_name"])
	// fmt.Println("email:", result["email"])
	// fmt.Println("age:", result["age"])
}
