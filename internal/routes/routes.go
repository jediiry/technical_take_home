package routes

import (
	"net/http"
	"technical_take_home/internal/handler"

	"github.com/gorilla/mux"
)

func RegisterRoutes(handler *handler.KeyValueHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.GetListKeys).Methods("GET")
	r.HandleFunc("/{key}", handler.Get).Methods("GET")
	r.HandleFunc("/{key}", handler.Put).Methods("PUT")
	r.HandleFunc("/{key}", handler.Delete).Methods("DELETE")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Route not found", http.StatusNotFound)
	})

	return r
}
