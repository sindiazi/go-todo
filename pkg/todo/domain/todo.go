package domain

import (
	"fmt"
	"math/rand"
	"reflect"
)

type TodoId struct {
	TodoUserId TodoUserId
	Id         int
}

func NewTodoId(userId TodoUserId) TodoId {
	return TodoId{TodoUserId: userId, Id: rand.Int() % 1000}
}

type TodoUserId struct {
	userId string
}

func NewTodoUserId(userId string) TodoUserId {
	return TodoUserId{userId: userId}
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
		todoMap[NewTodoId(todoUserId)] = todo
	}
	return &TodoList{
		TodoUserId: todoUserId,
		Todos:      todoMap,
	}
}

func (t *TodoList) Count() int {
	return len(t.Todos)
}

func (t *TodoList) Add(todo Todo) Todo {
	t.Todos[todo.ID] = todo
	return todo
}

func (t *TodoList) Remove(todoId TodoId) bool {
	_, todo := t.Todos[todoId]
	ok := reflect.ValueOf(todo).IsZero()

	if !ok {
		fmt.Println("Todo not found")
	} else {
		delete(t.Todos, todoId)
	}
	return ok
}

func (t *TodoList) Update(updatedTodo Todo) Todo {
	t.Todos[updatedTodo.ID] = updatedTodo
	return updatedTodo
}
func (t *TodoList) GetTodoById(todoId TodoId) Todo {
	return t.Todos[todoId]
}

func (t *TodoList) GetAllTodos() []Todo {
	values := make([]Todo, 0, len(t.Todos))
	for _, value := range t.Todos {
		values = append(values, value)
	}
	return values
}

//func (t *TodoList) PrintTodos(printer output.Printer) {
//	var todos []Todo
//	for _, todo := range t.Todos {
//		todos = append(todos, todo)
//	}
//	printer.Print(todos)
//}
