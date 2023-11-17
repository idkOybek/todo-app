package main

import (
	"os"

	"github.com/idkOybek/todo-app"
	"github.com/idkOybek/todo-app/pck/handler"
	"github.com/idkOybek/todo-app/pck/repository"
	"github.com/idkOybek/todo-app/pck/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConf(); err != nil {
		logrus.Fatalf("error with init conf : %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error with init env : %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error with connect to DB : %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error with init http server : %s", err.Error())
	}
}

func initConf() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
