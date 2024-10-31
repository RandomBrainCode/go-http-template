package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", server)

	fmt.Println("Listening on :8080...\nCtrl+C to exit")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
