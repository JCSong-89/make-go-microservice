package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":2323", &handlers.SearchHandler{})
	if err != nil {
		log.Fatal(err)
	}
}

