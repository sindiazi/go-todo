package usecases

import (
	"todos.com/m/v2/pkg/todo/ports/output"
	"todos.com/m/v2/pkg/todo/shared"
)

type DefaultAdminUseCases struct {
	todoRepository output.TodoListRepository
}

func NewAdminUseCases(todoRepository output.TodoListRepository) *DefaultAdminUseCases {
	return &DefaultAdminUseCases{todoRepository: todoRepository}
}

func (u *DefaultAdminUseCases) AddNewTodo(newTodo shared.TodoDto) shared.TodoDto {
	return u.todoRepository.SaveOrUpdate(newTodo)
}
func (u *DefaultAdminUseCases) RemoveTodo(todoId shared.TodoIdDto) bool {
	//return !reflect.ValueOf(u.todoRepository.GetByID(todoId)).IsZero()
	return false
}

func (u *DefaultAdminUseCases) UpdateTodo(newTodo shared.TodoDto) shared.TodoDto {
	return u.todoRepository.SaveOrUpdate(newTodo)
}

func (u *DefaultAdminUseCases) FindById(todoId shared.TodoIdDto) (shared.TodoDto, error) {
	return u.todoRepository.GetByID(todoId)
}
func (u *DefaultAdminUseCases) FindAllTodos(todoUserId string) []shared.TodoDto {
	return u.todoRepository.GetAllTodos(todoUserId)
}
