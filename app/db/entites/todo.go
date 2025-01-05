package entites

import "time"

type Todo struct {
	CreatedAt *time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"column:deletedAt"`
	Id        uint       `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Check     bool       `json:"check"`
}

type TodoDTO struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Check bool   `json:"check"`
}
