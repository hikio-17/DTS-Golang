package main

import (
	"assigment-1/helpers"
	"fmt"
	"os"
	"strconv"
)

var persons []helpers.Person = []helpers.Person{
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
	{
		No:           5,
		Nama:         "Asuna",
		Alamat:       "Jl. Merdeka No. 5",
		Pekerjaan:    "Freelance",
		AlasanGolang: "Iseng aja sih pengen nyoba hal baru aja sih",
	},
	{
		No:           4,
		Nama:         "Kirigaya Kazuto",
		Alamat:       "Jl. Merdeka No. 4",
		Pekerjaan:    "Freelance",
		AlasanGolang: "Iseng aja sih pengen nyoba hal baru aja sih",
	},
}

func main() {
	var arguments []string = os.Args

	if len(arguments) < 2 {
		fmt.Println("Silahkan masukkan id absen siswa")
		return
	}

	noAbsen, _ := strconv.Atoi(arguments[1])

	var person helpers.Person

	for _, p := range persons {
		if p.No == noAbsen {
			person = p
			break
		}
	}

	if person.Nama != "" {
		person.DisplayData()
	} else {
		fmt.Println("Data dengan no absen", noAbsen, "tidak ditemukan")
	}
}
