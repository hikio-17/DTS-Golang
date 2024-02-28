package main

import "fmt"

func main() {
	// checkCreateSim()
	// fmt.Printf(checkScore(8))
	// printNumber([]int{1, 2, 3})
	// printPola()

	// makeSlice()
	createMap()
}

func checkCreateSim() {
	var currentYear = 2021
	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim.")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}
}

func checkScore(score int) string {
	switch score {
	case 8:
		return "perfect"
	case 7:
		return "awesome"
	default:
		return "not bad"
	}
}

func printNumber(nums []int) {
	for i := 0; i < len(nums); i++ {
		fmt.Println("Angka", i)
	}
}

func printPola() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func makeSlice() {
	var fruits []string

	var fruits1 = []string{"apple", "banana", "manggo"}
	var fruits2 = []string{"durian", "stratfruit"}

	// menambahkan element didalam array
	// fruits = append(fruits, "apple", "banana", "manggo")

	fruits = append(fruits1, fruits2...)
	// nn := copy(fruits1, fruits2)
	// fmt.Printf("%#v", fruits)
	fmt.Println("%#v", fruits[1:4])
}

func createMap() {
	var person map[string]string // Deklarasi

	person = map[string]string{}

	person["name"] = "Fajri Muhammad Tio"
	person["age"] = "23"
	person["address"] = "Kota Bandung"

	// Or

	var person1 = map[string]string{
		"name":    "Fajri Muhammad Tio",
		"age":     "24",
		"address": "Kota Bandung",
	}
	_ = person1

	// perulangan pada map

	for key, value := range person {
		fmt.Println(key, ":", value)
	}

	// menghapus value dari sebuah map
	delete(person, "age")
	fmt.Println(person)

	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Valuen doesn't exist")
	}

	// Combine map & slice
	var people = []map[string]string{
		{"name": "Hikma Nelda", "age": "23"},
		{"name": "Fajri Muhammad Tio", "age": "24"},
	}

	fmt.Printf("people type %T", people)

	// fmt.Println(person1)
}
