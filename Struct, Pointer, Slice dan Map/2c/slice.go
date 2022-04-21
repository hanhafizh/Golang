package main

import "fmt"

func main() {
	var kursus = []string{"dbms", "server os", "networking", "web", "desktop", "erp"}
	var kursus_saya = kursus[1:5]

	kursus_saya = append(kursus_saya, "Manajemen Informatika")
	fmt.Println("Isi dari kursus adalah ", kursus)
	fmt.Println("Panjang Kursus ", len(kursus), "dan Kapasitas ", cap(kursus))

	fmt.Println("Isi dari kursus saya adalah", kursus_saya)
	fmt.Println("Panjang Kursus ", len(kursus_saya), "dan Kapasitas ", cap(kursus_saya))
}
