package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	No           int
	Nama         string
	Pekerjaan    string
	Alamat       string
	AlasanGolang string
}

var persons []Person = []Person{
	{
		No:           1,
		Nama:         "Fajri Muhammad Tio",
		Alamat:       "Jl. Merdeka No. 1",
		Pekerjaan:    "Mahasiswa",
		AlasanGolang: "Ingin belajar bahasa pemrograman baru yang powerful",
	},
	{
		No:           2,
		Nama:         "Hikma Nelda",
		Alamat:       "Jl. Merdeka No. 2",
		Pekerjaan:    "Tidak Bekerja",
		AlasanGolang: "Katanya over power untuk backend",
	},
	{
		No:           3,
		Nama:         "Kemku",
		Alamat:       "Jl. Merdeka No. 3",
		Pekerjaan:    "Freelance",
		AlasanGolang: "Iseng aja sih pengen nyoba hal baru aja sih",
	},
}

func main() {
	noAbsen, _ := strconv.Atoi(os.Args[1])

	var person Person

	for _, p := range persons {
		if p.No == noAbsen {
			person = p
			break
		}
	}

	if person.Nama != "" {
		fmt.Println("Nama: ", person.Nama)
		fmt.Println("Alamat: ", person.Alamat)
		fmt.Println("Pekerjaan: ", person.Pekerjaan)
		fmt.Println("Alasan Golang: ", person.AlasanGolang)
	} else {
		fmt.Println("Data dengan no absen", noAbsen, "tidak ditemukan")
	}

}
