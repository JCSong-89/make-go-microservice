package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	handler := newValidationHander(newHelloWorldHandler())
	http.Handle("/helloWorld", handler)

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationContextKey string

type validationHandler struct {
	next http.Handler
}

type helloWorldResponse struct {
	Message string  `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldHandler struct{}

func newValidationHander(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&request)

	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r  = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	respone := helloWorldResponse{Message: "Hello" + name }

	encoder := json.NewEncoder(rw)
	encoder.Encode(respone)
}

	