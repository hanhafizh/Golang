package main

import "fmt"

var i int

func main() {
	for j := 0; j <= 10; j++ {
		fmt.Print("Input nilai I : ")
		fmt.Scan(&i)

		if i <= 10 {
			if i%2 == 1 {
				fmt.Println(i, "Adalah angka Ganjil")
			} else {
				fmt.Println(i, "Adalah angka Geenap")
			}
		} else {
			fmt.Println("Pertanyaan selesai, karena I yang diinput sudah melebihi dari 10. Terimakasih")
			break
		}
	}
}
