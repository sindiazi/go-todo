package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"todos.com/m/v2/pkg/todo/adapters/output"
	"todos.com/m/v2/pkg/todo/usecases"
)

type TestStruct struct {
	Name string `json:"name"`
}

func main() {
	repo := output.NewInMemoryTodoRepository()
	adminUseCases := usecases.NewAdminUseCases(repo)
	todoWebAdminAdapter := output.NewTodoWebAdminAdapter(adminUseCases)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{username}", todoWebAdminAdapter.ListTodoHandler())
	r.Get("/{username}/{id}", todoWebAdminAdapter.FindTodoByIdHandler())

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("could not start server: %v", err)
	}
}
