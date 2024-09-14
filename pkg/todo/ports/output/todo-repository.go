package output

import "todos.com/m/v2/pkg/todo/domain"

type TodoListRepository interface {
	SaveOrUpdate(todo domain.TodoList) domain.TodoList
	GetAllTodos(userId domain.TodoUserId) []domain.Todo
	GetByID(todoId domain.TodoId) domain.Todo
	Delete(todoId domain.TodoId) domain.Todo
}
