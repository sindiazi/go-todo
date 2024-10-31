package output

import (
	"todos.com/m/v2/pkg/todo/domain"
	"todos.com/m/v2/pkg/todo/shared"
)

type TodoAdminPort interface {
	AddTodo(todo shared.TodoDto) shared.TodoDto
	RemoveTodo(todoId domain.TodoId) bool
	UpdateTodo(updatedTodo shared.TodoDto) shared.TodoDto
	ListTodos(todoUserId string) []shared.TodoDto
	FindTodoById(todoId shared.TodoIdDto) shared.TodoDto
}
