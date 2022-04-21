package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := 8060
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Printf("Server running on part %d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
