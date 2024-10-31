package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello from Go!"}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func main() {
	server := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", server)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Listening on :8080...\nCtrl+C to exit")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
