package main

import (
	"log"

	"github.com/idkOybek/todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error with init http server : %s", err.Error())
	}
}
