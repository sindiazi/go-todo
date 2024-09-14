package main

import (
	"todos.com/m/v2/pkg/todo/domain"
)

func main() {
	todos := []domain.Todo{
		{domain.NewTodoId(), "Create a new project", false, "Create a new project using Go"},
		{domain.NewTodoId(), "Create a another project", false, "Create a new project using Go"},
		{domain.NewTodoId(), "Create a yet another project", false, "Create a new project using Go"},
	}

	todolist := domain.NewPopulatedTodoList(todos)
	defaultPrinter := domain.DefaultPrinter{}
	todolist.PrintTodos(&defaultPrinter)

}
