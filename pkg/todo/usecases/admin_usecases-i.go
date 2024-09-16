package usecases

import "todos.com/m/v2/pkg/todo/domain"

type AdminUseCases interface {
	AddNewTodo(newTodo domain.Todo) domain.Todo
	RemoveTodo(todoId domain.TodoId) bool
	UpdateTodo(newTodo domain.Todo) domain.Todo
	FindById(todoId domain.TodoId) domain.Todo
	FindAllTodos(todoUserId domain.TodoUserId) []domain.Todo
}
