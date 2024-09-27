package shared

import (
	"todos.com/m/v2/pkg/todo/domain"
)

type TodoIdDto struct {
	ID         int    `json:"id"`
	TodoUserId string `json:"todoUserId"`
}

func (todoDto *TodoIdDto) FromTodoId(todo domain.TodoId) TodoIdDto {
	return TodoIdDto{
		ID:         todo.Id,
		TodoUserId: todo.TodoUserId.UserId,
	}
}

type TodoDto struct {
	TodoId      TodoIdDto `json:"todoId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
}

func (*TodoDto) NewTodoDto(userId string, todoId int, title string, description string, isCompleted bool) TodoDto {
	return TodoDto{
		TodoUserId: TodoIdDto{
			ID:         todoId,
			TodoUserId: userId,
		},
		Title:       title,
		Description: description,
		IsCompleted: isCompleted,
	}
}

func (todoDto *TodoDto) FromTodo(todo domain.Todo) TodoDto {
	todoIdDto := TodoIdDto{}

	return TodoDto{
		TodoUserId:  todoIdDto.FromTodoId(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
	}
}

type TodoListDto struct {
	Todos      []TodoDto `json:"todos"`
	TodoUserId string    `json:"todoUserId"`
}

func (*TodoListDto) FromTodoList(todoList domain.TodoList) TodoListDto {

	todos := []TodoDto{}

	for _, todo := range todoList.Todos {
		todoDto := TodoDto{}
		todos = append(todos, todoDto.FromTodo(todo))
	}

	return TodoListDto{
		TodoUserId: todoList.TodoUserId.UserId,
		Todos:      todos,
	}
}
