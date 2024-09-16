package output

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"todos.com/m/v2/pkg/todo/domain"
)

type MockAdminUseCases struct {
	mock.Mock
}

func (m *MockAdminUseCases) AddNewTodo(todo domain.Todo) domain.Todo {
	args := m.Called(todo)
	return args.Get(0).(domain.Todo)
}

func (m *MockAdminUseCases) RemoveTodo(todoId domain.TodoId) bool {
	args := m.Called(todoId)
	return args.Bool(0)
}

func (m *MockAdminUseCases) UpdateTodo(updatedTodo domain.Todo) domain.Todo {
	args := m.Called(updatedTodo)
	return args.Get(0).(domain.Todo)
}

func (m *MockAdminUseCases) FindAllTodos(todoUserId domain.TodoUserId) []domain.Todo {
	args := m.Called(todoUserId)
	return args.Get(0).([]domain.Todo)
}

func (m *MockAdminUseCases) FindById(todoId domain.TodoId) domain.Todo {
	args := m.Called(todoId)
	return args.Get(0).(domain.Todo)
}

func TestTodoWebAdminAdapter_AddTodo(t *testing.T) {
	mockUseCases := new(MockAdminUseCases)
	adapter := TodoWebAdminAdapter{adminUseCases: mockUseCases}
	todo := domain.Todo{ID: domain.NewTodoId(domain.NewTodoUserId("sindiazi")), Title: "Test Todo"}

	mockUseCases.On("AddNewTodo", todo).Return(todo)

	result := adapter.AddTodo(todo)

	assert.Equal(t, todo, result)
	mockUseCases.AssertExpectations(t)
}

func TestTodoWebAdminAdapter_RemoveTodo(t *testing.T) {
	mockUseCases := new(MockAdminUseCases)
	adapter := TodoWebAdminAdapter{adminUseCases: mockUseCases}
	todoId := domain.NewTodoId(domain.NewTodoUserId("sindiazi"))

	mockUseCases.On("RemoveTodo", todoId).Return(true)

	result := adapter.RemoveTodo(todoId)

	assert.True(t, result)
	mockUseCases.AssertExpectations(t)
}

func TestTodoWebAdminAdapter_UpdateTodo(t *testing.T) {
	mockUseCases := new(MockAdminUseCases)
	adapter := TodoWebAdminAdapter{adminUseCases: mockUseCases}
	updatedTodo := domain.Todo{ID: domain.NewTodoId(domain.NewTodoUserId("sindiazi")), Title: "Updated Todo"}

	mockUseCases.On("UpdateTodo", updatedTodo).Return(updatedTodo)

	result := adapter.UpdateTodo(updatedTodo)

	assert.Equal(t, updatedTodo, result)
	mockUseCases.AssertExpectations(t)
}

func TestTodoWebAdminAdapter_ListTodos(t *testing.T) {
	mockUseCases := new(MockAdminUseCases)
	adapter := TodoWebAdminAdapter{adminUseCases: mockUseCases}
	todoUserId := domain.NewTodoUserId("sindiazi")
	todos := []domain.Todo{
		domain.Todo{ID: domain.NewTodoId(domain.NewTodoUserId("sindiazi")), Title: "Updated Todo"},
	}

	mockUseCases.On("FindAllTodos", todoUserId).Return(todos)

	result := adapter.ListTodos(todoUserId)

	assert.Equal(t, todos, result)
	mockUseCases.AssertExpectations(t)
}

func TestTodoWebAdminAdapter_FindTodoById(t *testing.T) {
	mockUseCases := new(MockAdminUseCases)
	adapter := TodoWebAdminAdapter{adminUseCases: mockUseCases}
	todoId := domain.NewTodoId(domain.NewTodoUserId("sindiazi"))
	todo := domain.Todo{ID: todoId, Title: "Found Todo"}

	mockUseCases.On("FindById", todoId).Return(todo)

	result := adapter.FindTodoById(todoId)

	assert.Equal(t, todo, result)
	mockUseCases.AssertExpectations(t)
}
