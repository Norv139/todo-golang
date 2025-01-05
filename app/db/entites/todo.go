package entites

import "time"

type Todo struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Check     bool      `json:"check"`
}

type TodoDTO struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Check bool   `json:"check"`
}
