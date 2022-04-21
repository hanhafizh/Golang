package main

import "fmt"

func main() {
	var i int = 1
	var j int
	for {

		fmt.Print("Masukkan Angka Ke-", i, " = ")
		fmt.Scan(&j)
		i++

		if j%2 == 0 {
			fmt.Print(j, "adalah bilangan Genap\n\n")
		} else if j%2 != 0 {
			fmt.Print(j, "adalah bilangan Ganjil\n\n")
		} else {
			fmt.Print("Something Wrong!")
		}

		if i > 10 {
			fmt.Print("Pertanyaan selesai, karenakan angka yang di input sudah melebihi dari 10. Terimakasih")
			break
		}
	}
}
