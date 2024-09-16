package output

import "todos.com/m/v2/pkg/todo/domain"

type InMemoryTodoRepository struct {
	todoListsByUserIdMap map[domain.TodoUserId]domain.TodoList
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todoListsByUserIdMap: make(map[domain.TodoUserId]domain.TodoList),
	}
}

func (tl *InMemoryTodoRepository) SaveOrUpdate(todo domain.Todo) domain.Todo {
	targetTodoList := tl.todoListsByUserIdMap[todo.ID.TodoUserId]
	return targetTodoList.Add(todo)
}
func (tl *InMemoryTodoRepository) GetAllTodos(userId domain.TodoUserId) []domain.Todo {
	todoList := tl.todoListsByUserIdMap[userId]
	return todoList.GetAllTodos()

}
func (tl *InMemoryTodoRepository) GetByID(todoId domain.TodoId) domain.Todo {
	todoList := tl.todoListsByUserIdMap[todoId.TodoUserId]
	return todoList.GetTodoById(todoId)
}
func (tl *InMemoryTodoRepository) Delete(todoId domain.TodoId) bool {
	todoList := tl.todoListsByUserIdMap[todoId.TodoUserId]
	return todoList.Remove(todoId)
}
