package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homepage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "This is root route")
	fmt.Println("Endpoint hit: / ")
}
