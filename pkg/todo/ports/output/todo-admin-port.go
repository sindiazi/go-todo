package output

import "todos.com/m/v2/pkg/todo/domain"

type TodoAdminPort interface {
	AddTodo(todo domain.Todo) domain.Todo
	RemoveTodo(todoId domain.TodoId) bool
	UpdateTodo(updatedTodo domain.Todo) domain.Todo
	ListTodos() []domain.Todo
	FindTodoById(todoId domain.TodoId) domain.Todo
}
