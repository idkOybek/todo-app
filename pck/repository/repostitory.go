package repository

import (
	"github.com/idkOybek/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoLists interface {
}

type TodoItems interface {
}

type Repository struct {
	Authorization
	TodoLists
	TodoItems
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
