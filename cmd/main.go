package main

import (
	"log"
	"os"

	"github.com/dinachka/todo-app"
	"github.com/dinachka/todo-app/pkg/handler"
	"github.com/dinachka/todo-app/pkg/repository"
	"github.com/dinachka/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("initConfig", err.Error())
	}

		if err:= godotenv.Load(); err != nil {
			log.Fatalf("dotenv.Load", err.Error())
		}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username:  viper.GetString("db.username"),
		Password:  os.Getenv("DB_PASSWORD"),
		DBName:    viper.GetString("db.dbname"),
		SSLMode:   viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("dbErr", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
