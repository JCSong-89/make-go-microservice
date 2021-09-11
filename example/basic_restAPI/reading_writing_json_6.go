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

type helloWorldRequest struct {
	Name string `json:"name"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest // 받고자 하는 JSON 형식 지정

	decoder := json.NewDecoder(r.Body) // 넘어온 body 데이터 디코딩
	err := decoder.Decode(&request) // 형식이 맞는지 검사

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}
	// 지정한 형식에 맞게  응답데이터 만듬
	encoder := json.NewEncoder(w) // 응답객체 만듬
	encoder.Encode(response) // 위에서 만든 응답데이터 인코딩
	}

	