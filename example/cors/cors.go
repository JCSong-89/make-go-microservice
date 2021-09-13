package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	log.Fatal("Something")

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
// 해당 URI로 접근 시 OPSTIONS 요청인지 확인 후 아래의 Header를 추가하여 응답
	if r.Method == "OPTIONS" {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	//OPTIONS이 아닐 시 JSON 객체 리턴
	response := helloWorldResponse{Message: "Hello World"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}