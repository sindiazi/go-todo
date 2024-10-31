package usecases

import (
	"todos.com/m/v2/pkg/todo/shared"
)

type AdminUseCases interface {
	AddNewTodo(newTodo shared.TodoDto) shared.TodoDto
	RemoveTodo(todoId shared.TodoIdDto) bool
	UpdateTodo(newTodo shared.TodoDto) shared.TodoDto
	FindById(todoId shared.TodoIdDto) (shared.TodoDto, error)
	FindAllTodos(todoUserId string) []shared.TodoDto
}
