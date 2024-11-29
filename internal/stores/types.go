package stores

import "github.com/Gust4voSales/go-todo-app/internal/types"

type TodoStoreInterface interface {
	CreateTodo(todo types.Todo) error
	ListTodos() ([]types.Todo, error)
	GetTodo(id string) (*types.Todo, error)
	UpdateTodoCompleted(id string, status bool) error
	DeleteTodo(id string) error
}
