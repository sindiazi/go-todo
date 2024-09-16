package output

import "todos.com/m/v2/pkg/todo/domain"

type Printer interface {
	Print(todos []domain.Todo)
}

type DefaultPrinter struct {
}

func (p *DefaultPrinter) Print(todos []domain.Todo) {
	for _, todo := range todos {
		println(todo.Title, todo.Completed, todo.Description)
	}
	println("Total todos: ", len(todos))
}
