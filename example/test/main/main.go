package main

import (
	"log"
	"net/http"
	"test.com/handlers"
	"test.com/data"
)

func main() {
	handle := handlers.NewSearchHandler(&data.MockStore{})
	err := http.ListenAndServe(":2323", &handle)
	if err != nil {
		log.Fatal(err)
	}	
}

