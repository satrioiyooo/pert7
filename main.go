package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 9000 //4 digit npm terakhir

	http.Handle("/",http.FileServer((http.Dir("Polymer"))))
	fmt.Printf("server Running in http://localhost:%d",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port), nil))
}
