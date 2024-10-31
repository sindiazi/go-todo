package output

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"todos.com/m/v2/pkg/todo/shared"
	"todos.com/m/v2/pkg/todo/usecases"
)

type ErrorMessage struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
type TodoWebAdminAdapter struct {
	adminUseCases usecases.AdminUseCases
}

func NewTodoWebAdminAdapter(adminUseCases usecases.AdminUseCases) *TodoWebAdminAdapter {
	return &TodoWebAdminAdapter{adminUseCases: adminUseCases}
}

func (twa *TodoWebAdminAdapter) AddTodo(todo shared.TodoDto) shared.TodoDto {
	return twa.adminUseCases.AddNewTodo(todo)
}
func (twa *TodoWebAdminAdapter) RemoveTodo(todoId shared.TodoIdDto) bool {
	return twa.adminUseCases.RemoveTodo(todoId)
}
func (twa *TodoWebAdminAdapter) UpdateTodo(updatedTodo shared.TodoDto) shared.TodoDto {
	return twa.adminUseCases.UpdateTodo(updatedTodo)
}
func (twa *TodoWebAdminAdapter) ListTodos(todoUserId string) []shared.TodoDto {
	return twa.adminUseCases.FindAllTodos(todoUserId)
}
func (twa *TodoWebAdminAdapter) FindTodoById(todoId shared.TodoIdDto) (shared.TodoDto, error) {
	result, err := twa.adminUseCases.FindById(todoId)
	if err != nil {
		return shared.TodoDto{}, err
	} else {
		return result, nil

	}
}

func (twa *TodoWebAdminAdapter) ListTodoHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username := getPathIndex(r, 1)
		result, _ := json.Marshal(twa.ListTodos(username))
		println(result)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(result))
		if err != nil {
			log.Fatal("could not process the request", err)
		}
	}
}

func (twa *TodoWebAdminAdapter) FindTodoByIdHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username := getPathIndex(r, 1)
		id, _ := strconv.Atoi(getPathIndex(r, 2))
		todoId := shared.TodoIdDto{ID: id, TodoUserId: username}
		foundTodo, err := twa.FindTodoById(todoId)
		if err != nil {
			errorMessage := ErrorMessage{Error: err.Error(), Status: http.StatusNotFound}
			result, _ := json.Marshal(errorMessage)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_, writeErr := w.Write([]byte(result))
			if writeErr != nil {
				log.Fatal("could not process the request", writeErr)
			}
		} else {
			result, _ := json.Marshal(foundTodo)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte(result))
			if err != nil {
				log.Fatal("could not process the request", err)
			}
		}
	}
}

func getPathIndex(r *http.Request, index int) string {
	path := r.URL.Path
	return strings.Split(path, "/")[index]
}
