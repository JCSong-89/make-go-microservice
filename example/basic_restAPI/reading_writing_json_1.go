package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/hello-world", helloHandler)

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type helloWorldResponse struct {
	Message string  `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "helloWorld"}
	data, err := json.Marshal(response)
	
	if err != nil {
		panic("Oopps")
	}

	fmt.Fprint(w, string(data))
}

