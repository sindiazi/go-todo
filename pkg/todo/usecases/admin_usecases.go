package usecases

import (
	"reflect"
	"todos.com/m/v2/pkg/todo/domain"
	"todos.com/m/v2/pkg/todo/ports/output"
)

type DefaultAdminUseCases struct {
	todoRepository output.TodoListRepository
}

func NewAdminUseCases(todoRepository output.TodoListRepository) *DefaultAdminUseCases {
	return &DefaultAdminUseCases{todoRepository: todoRepository}
}

func (u *DefaultAdminUseCases) AddNewTodo(newTodo domain.Todo) domain.Todo {
	return u.todoRepository.SaveOrUpdate(newTodo)
}
func (u *DefaultAdminUseCases) RemoveTodo(todoId domain.TodoId) bool {
	return !reflect.ValueOf(u.todoRepository.GetByID(todoId)).IsZero()
}

func (u *DefaultAdminUseCases) UpdateTodo(newTodo domain.Todo) domain.Todo {
	return u.todoRepository.SaveOrUpdate(newTodo)
}

func (u *DefaultAdminUseCases) FindById(todoId domain.TodoId) domain.Todo {
	return u.todoRepository.GetByID(todoId)
}
func (u *DefaultAdminUseCases) FindAllTodos(todoUserId domain.TodoUserId) []domain.Todo {
	return u.todoRepository.GetAllTodos(todoUserId)
}
