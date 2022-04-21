package main

import "fmt"

func main() {

	var nUTS, nUAS int

	fmt.Println("Nilai UTS dengan bobot 30%")
	fmt.Println("Nilai UAS dengan bobot 70%")
	fmt.Println("Total nilai : UTS + UAS")
	fmt.Println()

	fmt.Println("Masukkan nilai UTS: ")
	fmt.Scan(&nUTS)
	fmt.Println("Masukkan nilai UAS: ")
	fmt.Scan(&nUAS)

	tUTS := nUTS * 30 / 100
	tUAS := nUAS * 70 / 100
	total := tUTS + tUAS

	fmt.Println()
	fmt.Println("Total nilai UTS dan UAS : ", total)

	if total <= 20 {
		fmt.Println("Anda mendapatkan grade : E")
	} else if total >= 21 && total <= 40 {
		fmt.Println("Anda mendapatkan grade : D")
	} else if total >= 41 && total <= 60 {
		fmt.Println("Anda mendapatkan grade : C")
	} else if total >= 61 && total <= 80 {
		fmt.Println("Anda mendapatkan grade : B")
	} else {
		fmt.Println("Anda mendapatkan grade : A")
	}

}
