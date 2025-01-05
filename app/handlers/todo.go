package todo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"main/db"
	"main/db/entites"
	"net/http"
	"strconv"
)

type handlerTodo struct {
	Store    map[int64]interface{}
	Router   *mux.Router
	Clients  *db.StoreClients
	Postgres *sql.DB
}

func CreateTodoRouter(
	r *mux.Router,
	sc *db.StoreClients,
) *mux.Router {

	s := handlerTodo{}
	s.Store = make(map[int64]interface{})
	s.Router = r
	s.Clients = sc

	var err interface{}
	if s.Postgres, err = sc.Postgres.DB(); err != nil {
		panic(err)
	}

	r.HandleFunc("/todo", s.createTodo).Methods("POST")
	r.HandleFunc("/todo/{pk}", s.updateTodo).Methods("PUT")
	r.HandleFunc("/todo/{pk}", s.deleteTodo).Methods("DELETE")

	r.HandleFunc("/todo/find", s.findTodo).Methods("GET")
	r.HandleFunc("/todo/get/{pk}", s.getTodo).Methods("GET")

	return r
}

func (s *handlerTodo) createTodo(w http.ResponseWriter, r *http.Request) {
	todoItem := entites.TodoDTO{}

	err := json.NewDecoder(r.Body).Decode(&todoItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created := entites.Todo{
		Name:  todoItem.Name,
		Desc:  todoItem.Desc,
		Check: todoItem.Check,
	}

	todoTable := s.Clients.Postgres.Table("todo")

	todoTable.Create(&created)
	todoTable.Where("id = ?", created.Id).First(&created) // подтягивает остальные данные типо CreateAt

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (s *handlerTodo) updateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["pk"])

	updated := entites.Todo{}
	todoTable := s.Clients.Postgres.Table("todo")

	err := json.NewDecoder(r.Body).Decode(&updated)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result := todoTable.First(&updated, id); result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.Clients.Postgres.Updates(&updated)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

func (s *handlerTodo) deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["pk"])

	deleteItem := entites.Todo{Id: uint(id)}

	todoTable := s.Clients.Postgres.Table("todo")

	if res := todoTable.First(&deleteItem); res.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	todoTable.Delete(deleteItem)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleteItem)
}

// TODO: доделать запрос так, чтобы можно было указывать парметры поиска
func (s *handlerTodo) findTodo(w http.ResponseWriter, r *http.Request) {

	var todoItems []entites.Todo

	todoTable := s.Clients.Postgres.Table("todo")
	fnFind := todoTable.Find

	if res := fnFind(&todoItems); res.Error != nil {
		fmt.Println(res.Error)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoItems)
}

func (s *handlerTodo) getTodo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["pk"])

	var todoItems []entites.Todo

	db := s.Clients.Postgres

	qb := db.Table("todo").Where("id = ?", id)

	fnFind := qb.Find

	if res := fnFind(&todoItems); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusBadRequest)
		return
	}

	if len(todoItems) == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todoItems[0])
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	}
	return
}
