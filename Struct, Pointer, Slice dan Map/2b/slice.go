package main

import "fmt"

func main() {
	var kursus = []string{"dbms", "server os", "networking", "web", "desktop", "erp"}
	var kursus_saya = kursus[1:5]

	fmt.Println(kursus_saya)
}
