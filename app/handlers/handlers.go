package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type createDTO struct {
	createdAt int64  `json:"CreatedAt"`
	updatedAt int64  `json:"updatedAt"`
	deletedAt int64  `json:"deletedAt"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
}

type HandlerTodo struct {
	Store map[int64]interface{}
}

func (s *HandlerTodo) Init() *HandlerTodo {
	s.Store = make(map[int64]interface{})

	return s
}

func (s *HandlerTodo) CreateTodo(w http.ResponseWriter, r *http.Request) {
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
func (s *HandlerTodo) UpdateTodo(w http.ResponseWriter, r *http.Request) {
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
func (s *HandlerTodo) DeleteTodo(w http.ResponseWriter, r *http.Request) {
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
func (s *HandlerTodo) FindTodo(w http.ResponseWriter, r *http.Request) {
	values := make([]interface{}, 0, len(s.Store))
	for _, v := range s.Store {
		values = append(values, v)
	}

	json.NewEncoder(w).Encode(values)
}

// TODO: доделать
func (s *HandlerTodo) GetTodo(w http.ResponseWriter, r *http.Request) {
	values := make([]interface{}, 0, len(s.Store))
	for _, v := range s.Store {
		values = append(values, v)
	}

	json.NewEncoder(w).Encode(values)
}
