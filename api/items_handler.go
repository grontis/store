package api

import (
	"encoding/json"
	"grontis/store/storage"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ItemsHandler struct {
	db storage.Storage
}

func AddItemRoutes(router *mux.Router, db storage.Storage) {
	itemsHandler := ItemsHandler{db: db}
	router.HandleFunc("/items", itemsHandler.getItems).Methods("GET")
	router.HandleFunc("/items/{id}", itemsHandler.getItem).Methods("GET")
	router.HandleFunc("/items", itemsHandler.createItem).Methods("POST")
	router.HandleFunc("/items/{id}", itemsHandler.updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", itemsHandler.deleteItem).Methods("DELETE")
}

func (i ItemsHandler) getItems(w http.ResponseWriter, r *http.Request) {
	items, err := i.db.GetItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(items)
	if err != nil {
		http.Error(w, "Error encoding json response", http.StatusInternalServerError)
		return
	}

	writeJsonResponse(w, jsonData)
}

func (i ItemsHandler) getItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId, ok := vars["id"]

	if !ok {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(itemId)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	item, err := i.db.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(item)
	if err != nil {
		http.Error(w, "Error encoding json response", http.StatusInternalServerError)
		return
	}
	writeJsonResponse(w, jsonData)
}

func (i *ItemsHandler) createItem(w http.ResponseWriter, r *http.Request) {
	var item storage.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid json body", http.StatusBadRequest)
		return
	}

	item, err = i.db.CreateItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (i *ItemsHandler) updateItem(w http.ResponseWriter, r *http.Request) {
	var item storage.Item
	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		http.Error(w, "Invalid json body", http.StatusBadRequest)
		return
	}

	item, err = i.db.UpdateItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (i *ItemsHandler) deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId, ok := vars["id"]

	if !ok {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(itemId)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	err = i.db.DeleteItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
