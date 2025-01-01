package todo

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"main/db"
	"net/http"
	"strconv"
	"time"
)

type createDTO struct {
	createdAt int64  `json:"CreatedAt"`
	updatedAt int64  `json:"updatedAt"`
	deletedAt int64  `json:"deletedAt"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
}

type handlerTodo struct {
	Store       map[int64]interface{}
	Router      *mux.Router
	Collections interface{}
}

func CreateTodoRouter(
	r *mux.Router,
	sc *db.StoreClients,
) *mux.Router {

	s := handlerTodo{}
	s.Store = make(map[int64]interface{})
	s.Router = r
	s.Collections = sc

	r.HandleFunc("/todo", s.createTodo).Methods("POST")
	r.HandleFunc("/todo/{pk}", s.updateTodo).Methods("PUT")
	r.HandleFunc("/todo", s.deleteTodo).Methods("DELETE")

	r.HandleFunc("/todo/find", s.findTodo).Methods("GET")
	r.HandleFunc("/todo/get/{pk}", s.getTodo).Methods("GET")

	return r
}

func (s *handlerTodo) createTodo(w http.ResponseWriter, r *http.Request) {
	id := len(s.Store) + 1

	updatedItem := createDTO{
		Id: id,
	}

	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	now := time.Now().Unix()
	updatedItem.createdAt = now

	s.Store[int64(id)] = updatedItem

	json.NewEncoder(w).Encode(updatedItem)
}

// TODO: доделать
func (s *handlerTodo) updateTodo(w http.ResponseWriter, r *http.Request) {
	//id := mux.Vars(r)["pk"]
	var updatedItem createDTO

	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	now := time.Now().Unix()
	updatedItem.createdAt = now
	s.Store[now] = updatedItem

	json.NewEncoder(w).Encode(updatedItem)
}

// TODO: доделать
func (s *handlerTodo) deleteTodo(w http.ResponseWriter, r *http.Request) {
	var updatedItem createDTO

	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	now := time.Now().Unix()
	updatedItem.createdAt = now
	s.Store[now] = updatedItem

	json.NewEncoder(w).Encode(updatedItem)
}

// TODO: доделать
func (s *handlerTodo) findTodo(w http.ResponseWriter, r *http.Request) {
	values := make([]interface{}, 0, len(s.Store))
	for _, v := range s.Store {
		values = append(values, v)
	}

	json.NewEncoder(w).Encode(values)
}

func (s *handlerTodo) getTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["pk"]

	i, _ := strconv.Atoi(id)

	values := s.Store[int64(i)]

	json.NewEncoder(w).Encode(values)
}
