package output

import (
	"encoding/json"
	"log"
	"net/http"
	"todos.com/m/v2/pkg/todo/domain"
	"todos.com/m/v2/pkg/todo/usecases"
)

type TodoWebAdminAdapter struct {
	adminUseCases usecases.AdminUseCases
}

func NewTodoWebAdminAdapter(adminUseCases usecases.AdminUseCases) *TodoWebAdminAdapter {
	return &TodoWebAdminAdapter{adminUseCases: adminUseCases}
}

func (twa *TodoWebAdminAdapter) AddTodo(todo domain.Todo) domain.Todo {
	return twa.adminUseCases.AddNewTodo(todo)
}
func (twa *TodoWebAdminAdapter) RemoveTodo(todoId domain.TodoId) bool {
	return twa.adminUseCases.RemoveTodo(todoId)
}
func (twa *TodoWebAdminAdapter) UpdateTodo(updatedTodo domain.Todo) domain.Todo {
	return twa.adminUseCases.UpdateTodo(updatedTodo)
}
func (twa *TodoWebAdminAdapter) ListTodos(todoUserId domain.TodoUserId) []domain.Todo {
	return twa.adminUseCases.FindAllTodos(todoUserId)
}
func (twa *TodoWebAdminAdapter) FindTodoById(todoId domain.TodoId) domain.Todo {
	return twa.adminUseCases.FindById(todoId)
}

func HandleAddTodo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//todoList := domain.NewTodoList(domain.NewTodoUserId("This is a test"))
		result, _ := json.Marshal(struct {
			test string
		}{test: "This is a test"})
		println("This is the printed line...")
		println(result)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(result))
		if err != nil {
			log.Fatal("could not process the request", err)
		}
	}
}
