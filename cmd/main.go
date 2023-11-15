package main

import (
	"log"

	"github.com/idkOybek/todo-app"
	"github.com/idkOybek/todo-app/pck/handler"
	"github.com/idkOybek/todo-app/pck/repository"
	"github.com/idkOybek/todo-app/pck/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConf(); err != nil {
		log.Fatalf("error with init conf : %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error with init http server : %s", err.Error())
	}
}

func initConf() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
