package main

import "fmt"

func main() {
	// Without type data
	// var name = "Fajri"
	// var age = 24

	//With tipe data
	// var name string = "Fajri"
	// var age int = 24

	//shortcut
	name := "Fajri"
	age := 24

	fmt.Println("nama saya adalah ==>", name)
	fmt.Println("Umur saya ==>", age)

	fmt.Printf("Nama saya adalah %s dan umur saya saat ini %d", name, age)
}
