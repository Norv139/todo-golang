package entites

type Todo struct {
	CreatedAt int64  `db:"createdAt"`
	UpdatedAt int64  `db:"updatedAt"`
	DeletedAt int64  `db:"deletedAt"`
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Desc      string `db:"desc"`
	Check     bool   `db:"check"`
}

type TodoDTO struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Check bool   `json:"check"`
}
