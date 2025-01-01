package todo

import (
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
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
	Store   map[int64]interface{}
	Router  *mux.Router
	Clients *db.StoreClients
}

func CreateTodoRouter(
	r *mux.Router,
	sc *db.StoreClients,
) *mux.Router {

	s := handlerTodo{}
	s.Store = make(map[int64]interface{})
	s.Router = r
	s.Clients = sc

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

	rows := s.Clients.Postgres.QueryRowContext(
		queryCtx,
		sqlInsert,
		createdItem.Desc,
		createdItem.Name,
	)
	rows.Scan(&createdEntity.Id)

	selectRows, _ := s.Clients.Postgres.Query(`select "id", "name" from todo`)

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

	log.Println(entitiesArr)

	json.NewEncoder(w).Encode(createdItem)
}

// TODO: доделать
func (s *handlerTodo) updateTodo(w http.ResponseWriter, r *http.Request) {
	//id := mux.Vars(r)["pk"]

	var updateDTO todoDTO
	//var todoEntity entites.Todo

	err := json.NewDecoder(r.Body).Decode(&updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(updateDTO)
}

// TODO: доделать
func (s *handlerTodo) deleteTodo(w http.ResponseWriter, r *http.Request) {
	var updatedItem todoDTO

	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	now := time.Now().Unix()
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
