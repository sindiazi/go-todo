package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"todos.com/m/v2/pkg/todo/adapters/output"
)

type TestStruct struct {
	Name string `json:"name"`
}

func main() {
	handleTodo := output.HandleAddTodo()
	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	r.Get("/", handleTodo)

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	result, _ := json.Marshal(TestStruct{
	//		name: "This is a test",
	//	})
	//	print()
	//	w.Header().Set("Content-Type", "application/json")
	//	w.Write([]byte(result))
	//})
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("could not start server: %v", err)
	}
}
