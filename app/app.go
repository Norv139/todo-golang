package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/db"
	todo "main/handlers"
	"net/http"
	"time"
)

func main() {

	storeConnections := db.InitConnections()
	router := mux.NewRouter()

	todo.CreateTodoRouter(router, storeConnections)

	// Start the HTTP server

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
