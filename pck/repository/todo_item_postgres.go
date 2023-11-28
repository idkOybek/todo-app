package repository

import (
	"github.com/idkOybek/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NowTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(userId, listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()

}
