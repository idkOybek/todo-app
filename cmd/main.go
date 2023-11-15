package main

import (
	"log"

	"github.com/idkOybek/todo-app"
	"github.com/idkOybek/todo-app/pck/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error with init http server : %s", err.Error())
	}
}
