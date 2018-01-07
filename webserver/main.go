package main

import (
	"fmt"
	"io"
	"net/http"
	"log"
)



func homePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>17Disney ETL2</h1>")
	reLaunch()
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":5000", nil)
	fmt.Println("17Disney ETL Service")
}