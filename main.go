package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", server)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
