package output

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todos.com/m/v2/pkg/todo/domain"
	"todos.com/m/v2/pkg/todo/shared"
)

func TestInMemoryTodoRepository_Delete(t *testing.T) {
	type fields struct {
		todoListsByUserIdMap map[string]shared.TodoListDto
	}
	type args struct {
		todoId shared.TodoIdDto
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "",
			fields: fields{
				todoListsByUserIdMap: nil,
			},
			args: args{
				todoId: shared.TodoIdDto{
					ID:         0,
					TodoUserId: "",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &InMemoryTodoRepository{
				todoListsByUserIdMap: tt.fields.todoListsByUserIdMap,
			}
			assert.Equalf(t, tt.want, tl.Delete(tt.args.todoId), "Delete(%v)", tt.args.todoId)
		})
	}
}

func TestInMemoryTodoRepository_GetAllTodos(t *testing.T) {
	type fields struct {
		todoListsByUserIdMap map[string]shared.TodoListDto
	}
	type args struct {
		userId domain.TodoUserId
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []domain.Todo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &InMemoryTodoRepository{
				todoListsByUserIdMap: tt.fields.todoListsByUserIdMap,
			}
			assert.Equalf(t, tt.want, tl.GetAllTodos(tt.args.userId.UserId), "GetAllTodos(%v)", tt.args.userId)
		})
	}
}

func TestInMemoryTodoRepository_GetByID(t *testing.T) {
	type fields struct {
		todoListsByUserIdMap map[string]shared.TodoListDto
	}
	type args struct {
		todoId shared.TodoIdDto
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   domain.Todo
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &InMemoryTodoRepository{
				todoListsByUserIdMap: tt.fields.todoListsByUserIdMap,
			}
			result, _ := tl.GetByID(tt.args.todoId)
			assert.Equalf(t, tt.want, result, "GetByID(%v)", tt.args.todoId)
		})
	}
}

func TestInMemoryTodoRepository_SaveOrUpdate(t *testing.T) {
	type fields struct {
		todoListsByUserIdMap map[string]shared.TodoListDto
	}
	type args struct {
		todo shared.TodoDto
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   domain.Todo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &InMemoryTodoRepository{
				todoListsByUserIdMap: tt.fields.todoListsByUserIdMap,
			}
			assert.Equalf(t, tt.want, tl.SaveOrUpdate(tt.args.todo), "SaveOrUpdate(%v)", tt.args.todo)
		})
	}
}

func TestNewInMemoryTodoRepository(t *testing.T) {
	tests := []struct {
		name        string
		want        *InMemoryTodoRepository
		expectError bool
	}{
		{
			" Some Create New InMemory TodoRepository",
			NewInMemoryTodoRepository(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewInMemoryTodoRepository(), "NewInMemoryTodoRepository() was tested")
		})
	}
}
