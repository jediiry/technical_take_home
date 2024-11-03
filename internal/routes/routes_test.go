package routes

import (
	"bytes"
	"technical_take_home/internal/database"
	"technical_take_home/internal/handler"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRoutes() *mux.Router {
	dataStore := database.NewDataStore()
	handler := handler.NewKeyValueHandler(dataStore)
	return RegisterRoutes(handler)
}

func TestRoutes(t *testing.T) {
	router := setupRoutes()

	tests := []struct {
		name     string
		method   string
		path     string
		body     []byte
		expected int
		parallel bool
	}{
		{"GET / returns 200", "GET", "/", nil, http.StatusOK, true},
		{"GET /:key returns 404 for unknown key", "GET", "/invalidKey", nil, http.StatusNotFound, true},
		{"PUT /:key updates a key", "PUT", "/new_key", []byte(`{"name": "new_value"}`), http.StatusOK, true},
		{"Invalid payload format", "PUT", "/new_key", []byte(`{"value": "new_value"}`), http.StatusBadRequest, true},
		{"DELETE /:key deletes a key", "DELETE", "/new_key", nil, http.StatusNotFound, true},
		{"No route returns 404", "GET", "/non_existent_route", nil, http.StatusNotFound, true},
		
		
		{"PUT /:key creates a new key", "PUT", "/new_key", []byte(`{"name": "new_value"}`), http.StatusOK, false},
		{"GET /:key return a 200", "GET", "/new_key", nil, http.StatusOK, false},
		{"DELETE /:key deletes a key", "DELETE", "/new_key", nil, http.StatusGone, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.parallel {
				t.Parallel()
			}
			req, err := http.NewRequest(test.method, test.path, bytes.NewBuffer(test.body))
			if err != nil {
				t.Fatal(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, test.expected, w.Code)
		})
	}
}
