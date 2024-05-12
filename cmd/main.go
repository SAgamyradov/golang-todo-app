package main

import (
	"log"

	todo "github.com/SAgamyradov/golang-todo-app"
	"github.com/SAgamyradov/golang-todo-app/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run(": 8000", handler.InitRoutes()); err != nil {
		log.Fatalf(": error occured while running http server: %s", err.Error())
	}
}
