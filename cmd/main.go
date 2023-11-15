package main

import (
	"log"

	"github.com/idkOybek/todo-app"
	"github.com/idkOybek/todo-app/pck/handler"
	"github.com/idkOybek/todo-app/pck/repository"
	"github.com/idkOybek/todo-app/pck/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error with init http server : %s", err.Error())
	}
}
