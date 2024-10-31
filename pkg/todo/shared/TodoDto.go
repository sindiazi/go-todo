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

func NewTodoDto(userId string, todoId int, title string, description string, isCompleted bool) TodoDto {
	return TodoDto{
		TodoId: TodoIdDto{
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
		TodoId:      todoIdDto.FromTodoId(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
	}
}

func (todoDto *TodoDto) ToTodo() domain.Todo {
	todoId := domain.NewExistingTodoId(domain.NewTodoUserId(todoDto.TodoId.TodoUserId), todoDto.TodoId.ID)
	return domain.NewTodo(todoId, todoDto.Title, todoDto.Description, todoDto.IsCompleted)

}

type TodoListDto struct {
	Todos      []TodoDto `json:"todos"`
	TodoUserId string    `json:"todoUserId"`
}

func (todoListDto *TodoListDto) FromTodoList(todoList domain.TodoList) TodoListDto {

	var todos []TodoDto
	for _, todo := range todoList.Todos {
		var todoDto = TodoDto{}
		todos = append(todos, todoDto.FromTodo(todo))
	}

	return TodoListDto{
		TodoUserId: todoList.TodoUserId.UserId,
		Todos:      todos,
	}
}

func (*TodoListDto) TodoTodoList(todoList TodoListDto) domain.TodoList {
	todoUserId := domain.NewTodoUserId(todoList.TodoUserId)
	var todos []domain.Todo
	for _, todo := range todoList.Todos {
		todos = append(todos, todo.ToTodo())
	}
	return domain.NewPopulatedTodoList(todoUserId, todos)
}
