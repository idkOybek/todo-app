package service

import (
	"github.com/idkOybek/todo-app"
	"github.com/idkOybek/todo-app/pck/repository"
)

type TodoListService struct {
	repo repository.TodoLists
}

func NewTodoListService(repo repository.TodoLists) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
