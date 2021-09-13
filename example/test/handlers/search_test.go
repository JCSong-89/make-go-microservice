package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"../data"
	_ "github.com/stretchr/testify/mock"
)
func TestSearchHandlerReturnBadRequestWhenNoSearchCriteriaIsSent(t *testing.T) {
	r, rw, handler := setupTest(&searchRequest{}, "GET", "/search", nil)


handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest{
		t.Errorf("Expected BadRequest %v", rw.Code)
	} // 400을 리턴받지 않았을 때 테스트 실패
}

func TestSearchHandlerReturnBadRequestWHenBlankSearchCriteriaIsSent(t *testing.T) {
		r, rw, handler := setupTest(&searchRequest{}, "POST", "/search", nil)

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest{
		t.Errorf("Expected BadRequest %v", rw.Code)
		} 
	}

type searchRequest struct {
	Query string `json:"query"`
} // vaildation 규격
type searchResponse struct {
	Kittens []data.Kitten `json:"kittens"`
}

func newSearchHandler(m *data.MockStore) searchHandler{
	return searchHandler{DataStore: m} // searchHandler 객체를 리턴
}

type searchHandler struct {
	DataStore *data.MockStore
} 

func (s *searchHandler)  ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	
	request := new(searchRequest)
	err := decoder.Decode(&request)
 
	if err != nil || len(request.Query) < 1 {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return 
	}

		kittens := s.DataStore.Search(request.Query)

		encoder := json.NewEncoder(rw)
		encoder.Encode(searchResponse{Kittens: kittens})
} // Serve HTTP http.hanlder 객체 validation 미통과시 400 리턴, Search 결과 encoding 리턴

func setupTest(d interface{}, m string, URI string, n io.Reader) (*http.Request, *httptest.ResponseRecorder, searchHandler) {
	h := newSearchHandler(&data.MockStore{})
	rw := httptest.NewRecorder()

	if d == nil {
		return httptest.NewRequest(m, URI, n), rw, h
	}

	body, _ := json.Marshal(d)
	return httptest.NewRequest(m, URI, bytes.NewReader(body)), rw, h
}
