package main

import (
	"log"

	"github.com/dinachka/todo-app"
	"github.com/dinachka/todo-app/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}
