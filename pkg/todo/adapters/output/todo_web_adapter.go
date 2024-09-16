package output

import (
	"todos.com/m/v2/pkg/todo/domain"
	"todos.com/m/v2/pkg/todo/usecases"
)

type TodoWebAdminAdapter struct {
	adminUseCases usecases.AdminUseCases
}

func (twa *TodoWebAdminAdapter) AddTodo(todo domain.Todo) domain.Todo {
	return twa.adminUseCases.AddNewTodo(todo)
}
func (twa *TodoWebAdminAdapter) RemoveTodo(todoId domain.TodoId) bool {
	return twa.adminUseCases.RemoveTodo(todoId)
}
func (twa *TodoWebAdminAdapter) UpdateTodo(updatedTodo domain.Todo) domain.Todo {
	return twa.adminUseCases.UpdateTodo(updatedTodo)
}
func (twa *TodoWebAdminAdapter) ListTodos(todoUserId domain.TodoUserId) []domain.Todo {
	return twa.adminUseCases.FindAllTodos(todoUserId)
}
func (twa *TodoWebAdminAdapter) FindTodoById(todoId domain.TodoId) domain.Todo {
	return twa.adminUseCases.FindById(todoId)
}
