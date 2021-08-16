package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBooks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":4,"title":"bola goes to school","Author":" werner","Year":9081},{"id":6,"title":"","Author":"","Year":0},{"id":5,"title":"","Author":"","Year":0},{"id":7,"title":"bola goes to school","Author":" werner","Year":5509}]`

	assert.Equal(t, expected, rr.Body.String())
}
