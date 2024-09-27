package output

import (
	"fmt"
	"todos.com/m/v2/pkg/todo/shared"
)

const CONSTANT string = "constant"

type InMemoryTodoRepository struct {
	todoListsByUserIdMap map[string]shared.TodoListDto
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todoListsByUserIdMap: make(map[string]shared.TodoListDto),
	}
}

func (tl *InMemoryTodoRepository) SaveOrUpdate(todo shared.TodoDto) shared.TodoDto {
	targetTodoList := tl.todoListsByUserIdMap[todo.TodoId.TodoUserId]
	todos := append(targetTodoList.Todos, todo)
	_ = fmt.Sprintf("Added todo to List: %d", len(todos))
	return todo
}

func (tl *InMemoryTodoRepository) GetAllTodos(userId string) []shared.TodoDto {
	todoList := tl.todoListsByUserIdMap[userId]
	return todoList.Todos

}
func (tl *InMemoryTodoRepository) GetByID(todoIdToFind shared.TodoIdDto) (shared.TodoDto, error) {
	todoList := tl.todoListsByUserIdMap[todoIdToFind.TodoUserId]
	for _, todoId := range todoList.Todos {
		if todoId.TodoId.ID == todoIdToFind.ID {
			return todoId, nil
		}
	}
	return shared.TodoDto{}, fmt.Errorf("todo with id %d not found", todoIdToFind.ID)
}
func (tl *InMemoryTodoRepository) Delete(todoId shared.TodoIdDto) bool {
	todoList := tl.todoListsByUserIdMap[todoId.TodoUserId]
	return todoList.Remove(todoId)
}
