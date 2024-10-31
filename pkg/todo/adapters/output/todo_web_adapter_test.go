package output

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"todos.com/m/v2/pkg/todo/shared"
)

type MockAdminUseCases struct {
	mock.Mock
}

func (m *MockAdminUseCases) AddNewTodo(todo shared.TodoDto) shared.TodoDto {
	args := m.Called(todo)
	return args.Get(0).(shared.TodoDto)
}

func (m *MockAdminUseCases) RemoveTodo(todoId shared.TodoIdDto) bool {
	args := m.Called(todoId)
	return args.Bool(0)
}

func (m *MockAdminUseCases) UpdateTodo(updatedTodo shared.TodoDto) shared.TodoDto {
	args := m.Called(updatedTodo)
	return args.Get(0).(shared.TodoDto)
}

func (m *MockAdminUseCases) FindAllTodos(todoUserId string) []shared.TodoDto {
	args := m.Called(todoUserId)
	return args.Get(0).([]shared.TodoDto)
}

func (m *MockAdminUseCases) FindById(todoId shared.TodoIdDto) (shared.TodoDto, error) {
	args := m.Called(todoId)
	return args.Get(0).(shared.TodoDto), nil
}

func TestTodoWebAdminAdapter_AddTodo(t *testing.T) {
	mockUseCases := new(MockAdminUseCases)
	adapter := TodoWebAdminAdapter{adminUseCases: mockUseCases}
	todo := shared.NewTodoDto("sindiazi", 0, "Test Todo", "Test Description", false)

	mockUseCases.On("AddNewTodo", todo).Return(todo)

	result := adapter.AddTodo(todo)

	assert.Equal(t, todo, result)
	mockUseCases.AssertExpectations(t)
}

func TestTodoWebAdminAdapter_RemoveTodo(t *testing.T) {
	mockUseCases := new(MockAdminUseCases)
	adapter := TodoWebAdminAdapter{adminUseCases: mockUseCases}
	todoId := shared.TodoIdDto{ID: 0, TodoUserId: "sindiazi"}

	mockUseCases.On("RemoveTodo", todoId).Return(true)

	result := adapter.RemoveTodo(todoId)

	assert.True(t, result)
	mockUseCases.AssertExpectations(t)
}

func TestTodoWebAdminAdapter_UpdateTodo(t *testing.T) {
	mockUseCases := new(MockAdminUseCases)
	adapter := TodoWebAdminAdapter{adminUseCases: mockUseCases}
	updatedTodo := shared.NewTodoDto("sindiazi", 0, "Updated Test Todo", "Test Description", false)

	mockUseCases.On("UpdateTodo", updatedTodo).Return(updatedTodo)

	result := adapter.UpdateTodo(updatedTodo)

	assert.Equal(t, updatedTodo, result)
	mockUseCases.AssertExpectations(t)
}

func TestTodoWebAdminAdapter_ListTodos(t *testing.T) {
	mockUseCases := new(MockAdminUseCases)
	adapter := TodoWebAdminAdapter{adminUseCases: mockUseCases}
	updatedTodo := shared.NewTodoDto("sindiazi", 0, "Updated Test Todo", "Test Description", false)
	todos := []shared.TodoDto{
		updatedTodo,
	}

	mockUseCases.On("FindAllTodos", updatedTodo.TodoId.TodoUserId).Return(todos)

	result := adapter.ListTodos(updatedTodo.TodoId.TodoUserId)

	assert.Equal(t, todos, result)
	mockUseCases.AssertExpectations(t)
}

func TestTodoWebAdminAdapter_FindTodoById(t *testing.T) {
	mockUseCases := new(MockAdminUseCases)
	adapter := TodoWebAdminAdapter{adminUseCases: mockUseCases}
	todo := shared.NewTodoDto("sindiazi", 0, "Updated Test Todo", "Test Description", false)

	mockUseCases.On("FindById", todo.TodoId).Return(todo)

	result := adapter.FindTodoById(todo.TodoId)

	assert.Equal(t, todo, result)
	mockUseCases.AssertExpectations(t)
}
