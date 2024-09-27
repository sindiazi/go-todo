package output

import (
	"todos.com/m/v2/pkg/todo/shared"
)

type TodoListRepository interface {
	SaveOrUpdate(todo shared.TodoDto) shared.TodoDto
	GetAllTodos(userId string) []shared.TodoDto
	GetByID(todoId shared.TodoIdDto) shared.TodoDto
	Delete(todoId shared.TodoIdDto) shared.TodoDto
}
