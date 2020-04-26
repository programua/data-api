package test

import (
	"data_api.com/data_api/handler"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTopHandler(t *testing.T) {
	router := handler.MakeRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/top", nil)
	router.ServeHTTP(w, req)

	var tr map[string]string //TopResponse

	if err := json.NewDecoder(w.Body).Decode(&tr); err != nil {
		t.Error("can't decode response message")
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "top ok gin!", tr["message"])
}

func TestListHandler(t *testing.T) {
	router := handler.MakeRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/list", nil)
	router.ServeHTTP(w, req)

	var lr map[string]string //ListResponse

	if err := json.NewDecoder(w.Body).Decode(&lr); err != nil {
		t.Error("can't decode response message")
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "town_list gin ok!", lr["message"])
}
