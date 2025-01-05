package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/db"
	todo "main/handlers"
	"main/utils/middleware"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	storeConnections := db.InitConnections()
	router := mux.NewRouter()

	todo.CreateTodoRouter(router, storeConnections)

	// Start the HTTP server

	logMiddleware := middleware.NewLogMiddleware(logger)
	router.Use(logMiddleware.Func())

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
