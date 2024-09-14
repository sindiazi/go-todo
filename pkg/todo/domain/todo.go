package domain

import "math/rand"

type Printer interface {
	Print(todos []Todo)
}

type DefaultPrinter struct {
}

func (p *DefaultPrinter) Print(todos []Todo) {
	for _, todo := range todos {
		println(todo.ID.id, todo.Title, todo.Completed, todo.Description)
	}
	println("Total todos: ", len(todos))
}

type TodoId struct {
	todoUserId TodoUserId
	id         int
}

func NewTodoId() TodoId {
	return TodoId{id: rand.Int() % 1000}
}

type TodoUserId struct {
	id string
}

func NewTodoUserId(userId string) TodoUserId {
	return TodoUserId{id: userId}
}

type Todo struct {
	ID          TodoId
	Title       string
	Completed   bool
	Description string
}

type TodoList struct {
	TodoUserId TodoUserId
	Todos      map[TodoId]Todo
}

func NewTodoList(todoUserId TodoUserId) *TodoList {
	return NewPopulatedTodoList(todoUserId, []Todo{})
}

func NewPopulatedTodoList(todoUserId TodoUserId, todos []Todo) *TodoList {
	todoMap := make(map[TodoId]Todo)
	for _, todo := range todos {
		todoMap[NewTodoId()] = todo
	}
	return &TodoList{
		TodoUserId: todoUserId,
		Todos:      todoMap,
	}
}

func (t *TodoList) Count() int {
	return len(t.Todos)
}

func (t *TodoList) Add(todo Todo) {
	todoId := NewTodoId()
	t.Todos[todoId] = todo
}

func (t *TodoList) Remove(todoId TodoId) {
	delete(t.Todos, todoId)
}

func (t *TodoList) Update(updatedTodo Todo) Todo {
	t.Todos[updatedTodo.ID] = updatedTodo
	return updatedTodo
}
func (t *TodoList) GetTodoById(todoId TodoId) Todo {
	return t.Todos[todoId]
}

func (t *TodoList) PrintTodos(printer Printer) {
	var todos []Todo
	for _, todo := range t.Todos {
		todos = append(todos, todo)
	}
	printer.Print(todos)
}
