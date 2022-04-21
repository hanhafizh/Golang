package main

import "fmt"

func main() {
	var uts, uas, total_nilai float64
	fmt.Print("Masukkan Nilai UTS : ")
	fmt.Scan(&uts)
	fmt.Print("\nMasukkan Nilai UAS: ")
	fmt.Scan(&uas)

	fmt.Print("\nNilai UTS = ", uts)
	fmt.Print("\nNilai UTS = ", uas)

	total_nilai = (uts * 0.3) + (uas * 0.7)
	fmt.Print("\nTotal Nilai = ", total_nilai)

	if total_nilai <= 20 {
		fmt.Print("\nGrade E")
	} else if total_nilai <= 40 {
		fmt.Print("\nGrade D")
	} else if total_nilai <= 60 {
		fmt.Print("\nGrade C")
	} else if total_nilai <= 80 {
		fmt.Print("\nGrade B")
	} else {
		fmt.Print("\nGrade A")
	}
}
