package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr, store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// register services
	userService := NewUserService(s.store)
	userService.RegisterRoutes(subRouter)

	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(subRouter)

	log.Println("Starting API server at ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subRouter))
}
