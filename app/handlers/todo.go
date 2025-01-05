package todo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"main/db"
	"main/db/entites"
	"main/utils"
	"net/http"
	"strconv"
	"time"
)

type todoDTO struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Check bool   `json:"check"`
}

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
	r.HandleFunc("/todo", s.deleteTodo).Methods("DELETE")

	r.HandleFunc("/todo/find", s.findTodo).Methods("GET")
	r.HandleFunc("/todo/get/{pk}", s.getTodo).Methods("GET")

	return r
}

func (s *handlerTodo) createTodo(w http.ResponseWriter, r *http.Request) {
	createdItem := todoDTO{}

	err := json.NewDecoder(r.Body).Decode(&createdItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdEntity := entites.Todo{
		Name: createdItem.Name,
		Desc: createdItem.Desc,
	}

	queryCtx, queryCtxFm := utils.GetCtx()
	defer queryCtxFm()

	sqlInsert := `
	insert into
    todo ("id", "desc", "name")
	values (default, $1, $2)
	RETURNING "id"
	`

	rows := s.Postgres.QueryRowContext(
		queryCtx,
		sqlInsert,
		createdItem.Desc,
		createdItem.Name,
	)
	rows.Scan(&createdEntity.Id)

	selectRows, _ := s.Postgres.Query(`select "id", "name" from todo`)

	defer selectRows.Close()

	var entitiesArr []interface{}
	for selectRows.Next() {
		var (
			id   int64
			name string
		)
		rows.Scan(&id, &name)
		entitiesArr = append(entitiesArr, name)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdItem)
}

// TODO: ничего не делает - доделать
func (s *handlerTodo) updateTodo(w http.ResponseWriter, r *http.Request) {
	//id := mux.Vars(r)["pk"]

	var updateDTO todoDTO
	//var todoEntity entites.Todo

	err := json.NewDecoder(r.Body).Decode(&updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateDTO)
}

// TODO: ничего не делает - доделать
func (s *handlerTodo) deleteTodo(w http.ResponseWriter, r *http.Request) {
	var updatedItem todoDTO

	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	now := time.Now().Unix()
	s.Store[now] = updatedItem

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedItem)
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
