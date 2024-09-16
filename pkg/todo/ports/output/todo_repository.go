package output

import "todos.com/m/v2/pkg/todo/domain"

type TodoListRepository interface {
	SaveOrUpdate(todo domain.Todo) domain.Todo
	GetAllTodos(userId domain.TodoUserId) []domain.Todo
	GetByID(todoId domain.TodoId) domain.Todo
	Delete(todoId domain.TodoId) domain.Todo
}
