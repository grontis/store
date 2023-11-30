package api

import (
	"encoding/json"
	"fmt"
	"grontis/store/models"
	"grontis/store/storage"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	db storage.Storage
}

func AddAuthRoutes(router *mux.Router, db storage.Storage) {
	authHandler := AuthHandler{db: db}
	router.HandleFunc("/auth/users", authHandler.getUsers).Methods("GET")
	router.HandleFunc("/auth/users/{id}", authHandler.getUser).Methods("GET")
	router.HandleFunc("/auth/users", authHandler.createUser).Methods("POST")
	router.HandleFunc("/auth/users", authHandler.updateUser).Methods("PUT")
	router.HandleFunc("/auth/users/{id}", authHandler.deleteUser).Methods("DELETE")

	router.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
}

func (a AuthHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.db.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error encoding json response", http.StatusInternalServerError)
		return
	}

	writeJsonResponse(w, jsonData)
}

func (a AuthHandler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, ok := vars["id"]

	if !ok {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(userId)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	user, err := a.db.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding json response", http.StatusInternalServerError)
		return
	}

	writeJsonResponse(w, jsonData)
}

func (a AuthHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err = a.db.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding json response", http.StatusInternalServerError)
		return
	}

	writeJsonResponse(w, jsonData)
}

func (a AuthHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err = a.db.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding json response", http.StatusInternalServerError)
		return
	}

	writeJsonResponse(w, jsonData)
}

func (a AuthHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, ok := vars["id"]

	if !ok {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(userId)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	err = a.db.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.db.GetUserByUsername(credentials.Username)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid credentials"), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
