package usecases

import (
	"reflect"
	"todos.com/m/v2/pkg/todo/domain"
	"todos.com/m/v2/pkg/todo/ports/output"
)

type AdminUseCases struct {
	todoRepository output.TodoListRepository
}

func (u *AdminUseCases) AddNewTodo(newTodo domain.Todo) domain.Todo {
	return u.todoRepository.SaveOrUpdate(newTodo)
}
func (u *AdminUseCases) RemoveTodo(todoId domain.TodoId) bool {
	return !reflect.ValueOf(u.todoRepository.GetByID(todoId)).IsZero()
}

func (u *AdminUseCases) UpdateTodo(newTodo domain.Todo) domain.Todo {
	return u.todoRepository.SaveOrUpdate(newTodo)
}

func (u *AdminUseCases) FindById(todoId domain.TodoId) domain.Todo {
	return u.todoRepository.GetByID(todoId)
}
func (u *AdminUseCases) FindAllTodos() []domain.Todo {
	return u.todoRepository.GetAll()
}
