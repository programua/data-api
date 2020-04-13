package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newHttpClient() *http.Client {
	return new(http.Client)
}

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/top", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandleFunc
}