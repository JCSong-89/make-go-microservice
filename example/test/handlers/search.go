package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"test.com/data"
)

type searchRequest struct {
	Query string `json:"query"`
} // vaildation 규격
type searchResponse struct {
	Kittens []data.Kitten `json:"kittens"`
}


func NewSearchHandler(m *data.MockStore) SearchHandler{
	return SearchHandler{DataStore: m} // searchHandler 객체를 리턴
}

type SearchHandler struct {
	DataStore *data.MockStore
} 

func (s *SearchHandler)  ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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

func setupTest(d interface{}, m string, URI string, n io.Reader) (*http.Request, *httptest.ResponseRecorder, SearchHandler) {
	h := NewSearchHandler(&data.MockStore{})
	rw := httptest.NewRecorder()

	if d == nil {
		return httptest.NewRequest(m, URI, n), rw, h
	}

	body, _ := json.Marshal(d)
	return httptest.NewRequest(m, URI, bytes.NewReader(body)), rw, h
}
