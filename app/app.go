package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/handlers"
	"net/http"
)

func main() {

	hTodo := handlers.HandlerTodo{}

	router := mux.NewRouter()

	router.HandleFunc("/todo", hTodo.CreateTodo).Methods("POST")
	router.HandleFunc("/todo", hTodo.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo", hTodo.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todo", hTodo.FindTodo).Methods("GET")

	router.HandleFunc("/todo/find", hTodo.FindTodo).Methods("POST")
	router.HandleFunc("/todo/get", hTodo.GetTodo).Methods("POST")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":3000", router))
}
