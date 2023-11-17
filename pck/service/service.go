package service

import (
	"github.com/idkOybek/todo-app"
	"github.com/idkOybek/todo-app/pck/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoLists interface {
}

type TodoItems interface {
}

type Service struct {
	Authorization
	TodoLists
	TodoItems
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
