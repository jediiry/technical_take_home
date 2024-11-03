package handler

import (
	"encoding/json"
	"net/http"
	"technical_take_home/internal/database"

	"github.com/gorilla/mux"
)

type RequestBody struct {
	Name string `json:"name"`
}

type KeyValueHandler struct {
	store *database.DataStore
}

func NewKeyValueHandler(store *database.DataStore) *KeyValueHandler {
	return &KeyValueHandler{
		store: store,
	}
}

func (h *KeyValueHandler) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["key"]

	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	value, exists := h.store.Get(key)
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": value})
}

func (h *KeyValueHandler) GetListKeys(w http.ResponseWriter, r *http.Request) {
	keys := h.store.GetListKeys()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]string{"result": keys})
}

func (h *KeyValueHandler) Put(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["key"]

	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if requestBody.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	h.store.Put(key, string(requestBody.Name))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": "ok"})
}

func (h *KeyValueHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	isDeleted := h.store.Delete(key)
	if !isDeleted {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusGone)
	json.NewEncoder(w).Encode(map[string]string{"result": "Item has been deleted"})
}
