package handlers

import (
	"net/http"
	"testing"
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
