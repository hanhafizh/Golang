package main

import "fmt"

var nilai1, nilai2, sisa_bagi int

func main() {
	defer fmt.Println("---- SELESAI ----")

	fmt.Print("Masukkan Bilangan 1: ")
	fmt.Scan(&nilai1)

	fmt.Print("Masukkan Bilangan 2: ")
	fmt.Scan(&nilai2)

	hasil := nilai1 / nilai2
	sisa_bagi = nilai1 % nilai2

	fmt.Println("Hasil dari Nilai1/Nilai2 = ", hasil)
	fmt.Println("Sisa bagi dari Nilai 1 / Nilai 2 = ", sisa_bagi)
}
