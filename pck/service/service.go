package service

import (
	"github.com/idkOybek/todo-app/pck/repository"
)

type Authorization interface {
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
	return &Service{}
}
