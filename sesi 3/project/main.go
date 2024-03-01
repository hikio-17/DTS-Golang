package main

import (
	"fmt"
	"math"
	"project/helpers"
	"strings"
)

func main() {
	// greet("Fajri", 24)
	// fmt.Printf(greeting("Hello", []string{"Fajir", "Hikma"}))

	// var diameter float64 = 23
	// fmt.Println(calculated(float64(diameter)))

	// createCloseure()
	// createIIFE()

	var person = helpers.Person{}

	helpers.Greet()
	person.Invokegreet()
	// helpers.greet()

}

// func greet (name, address string)
func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old.", name, age)
}

func greeting(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

func calculated(d float64) (area float64, circumReference float64) {
	area = math.Phi * math.Pow(d/2, 2)
	circumReference = math.Phi * d

	return
}

func createCloseure() {
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(evenNumbers(numbers...))
}

func createIIFE() {
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("false")

	fmt.Println(isPalindrome)
}
