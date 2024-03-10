package helpers

import "fmt"

func (p Person) DisplayData() {
	fmt.Println("Nama: ", p.Nama)
	fmt.Println("Alamat: ", p.Alamat)
	fmt.Println("Pekerjaan: ", p.Pekerjaan)
	fmt.Println("Alasan Golang: ", p.AlasanGolang)
}
