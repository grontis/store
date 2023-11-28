package api

import (
	"fmt"
	"grontis/store/storage"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	db     storage.Storage
	router mux.Router
}

func NewServer(db storage.Storage, router *mux.Router) *Server {
	s := &Server{
		db:     db,
		router: *router,
	}

	AddItemRoutes(&s.router, s.db)
	return s
}

func (s Server) ListenAndServe(port string) {
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, &s.router)
}

func writeJsonResponse(w http.ResponseWriter, jsonData []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
