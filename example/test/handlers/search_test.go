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
	} // 400을 리턴받지 않았을 때 테스트 실패
}

type searchRequest struct {
	Name string `json:"name"`
} // vaildation 규격

func newSearchHandler() searchHandler{
	return searchHandler{} // searchHandler 객체를 리턴
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
} // Serve HTTP http.hanlder 객체 validation 미통과시 400 리턴
