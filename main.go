package main

import (
	"grontis/store/api"
	"grontis/store/storage"

	"github.com/gorilla/mux"
)

func main() {
	server := api.NewServer(storage.NewMemory(), mux.NewRouter())
	server.ListenAndServe("5000")
}
