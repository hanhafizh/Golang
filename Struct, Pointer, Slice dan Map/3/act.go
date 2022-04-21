package main

import "fmt"

func main() {
	var data = map[string]mahasiswa{
		"50420460": {
			"Farhan Hafizh",
			"2IA05",
		},
		"50420470": {
			"Aceng Toyib",
			"2IA05",
		},
	}
	var search string
	fmt.Print("Masukkan NPM anda? ")
	fmt.Scanf("%s", &search)

	var value, ok = data[search]

	if ok {
		fmt.Printf("Nama anda %s \nKelas Anda %s", value.Nama, value.Kelas)
	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}

type mahasiswa struct {
	Nama  string
	Kelas string
}
