package main

import "fmt"

const pi = 22 / 7

var jari int

func main() {
	fmt.Print("Masukkan Jari - Jari Lingkaran = ")
	fmt.Scan(&jari)
	luas := pi * (jari * jari)
	fmt.Println("Luar Lingkaran = ", luas)
}
