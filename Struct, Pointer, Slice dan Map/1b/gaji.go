package main

import "fmt"

func main() {
	var gajiSekarang, ekspektasiGaji int
	fmt.Print("Masukkan gaji Anda : ")
	fmt.Scan(&gajiSekarang)

	fmt.Print("Masukkan gaji yang anda inginkan : ")
	fmt.Scan(&ekspektasiGaji)

	naikanGaji(gajiSekarang, ekspektasiGaji)
	fmt.Printf("\nGaji Anda Sekarang %d\n", gajiSekarang)
}
func naikanGaji(gajiAwal int, gajiAkhir int) {
	gajiAwal = gajiAkhir
}
