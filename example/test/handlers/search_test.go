package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)
func TestSearchHanlderReturnBadRequestWhenNoSearchCriteriaIsSent(t *testing.T) {
	handler := newSearchHandler()
	request := httptest.NewRequest("GET", "/search", nil)
	response := httptest.NewRecorder()

handler.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest{
		t.Errorf("Expected BadRequest %v", response.Code)
	}
}

type searchRequest struct {
	Name string `json:"name"`
}

func newSearchHandler() searchHandler{
	return searchHandler{}
}

type searchHandler struct {
}

func (s *searchHandler)  ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	
	request := new(searchRequest)
	err := decoder.Decode(&request)

	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
}