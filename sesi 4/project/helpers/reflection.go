package helpers

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func Reflection() {
	// var number = 20
	// var reflectValue = reflect.ValueOf(number)

	// fmt.Println("tipe variabel : ", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variabel :", reflectValue.Int())
	// }

	// // Jika nilai hanya diperlukan untuk ditampikan ke output, bisa menggunakan Interface(), melalui method ini segala jenis nilai bisa diakases dengan mudah.
	// fmt.Println("nilai variabel dengan method Interface : ", reflectValue.Interface().(int))

	// example
	// informasi mengenai method juga bisa diakses lewat reflect, syaratnya masih sama seperti pada pengaksesan property, yaitu harus bermodifier public.

	var s1 = &student{Name: "John Wick", Grade: 2}
	fmt.Println("Nama : ", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")

	method.Call([]reflect.Value{
		reflect.ValueOf("Wick"),
	})

	fmt.Println("Nama : ", s1.Name)

	s1.SetName("Fajri")

	fmt.Println("Nama : ", s1.Name)
}
