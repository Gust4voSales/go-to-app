package services

import (
	"fmt"

	storeTypes "github.com/Gust4voSales/go-todo-app/internal/stores"
	"github.com/Gust4voSales/go-todo-app/internal/types"
	"github.com/google/uuid"
)

type TodoService struct {
	store storeTypes.TodoStoreInterface
}

func NewTodoService(store storeTypes.TodoStoreInterface) *TodoService {
	return &TodoService{
		store: store,
	}
}

func (ts *TodoService) ListTodos() ([]types.Todo, error) {
	todos, err := ts.store.ListTodos()

	if err != nil {
		return []types.Todo{}, err
	}

	return todos, nil
}

func (ts *TodoService) GetTodo(id string) (*types.Todo, error) {
	todo, err := ts.store.GetTodo(id)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (ts *TodoService) CreateTodo(content string) (*types.Todo, error) {
	ID := uuid.New()

	todo := types.Todo{
		ID:        ID.String(),
		Content:   content,
		Completed: false,
	}

	if err := ts.store.CreateTodo(todo); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (ts *TodoService) ToggleTodoCompleted(id string) (*types.Todo, error) {
	todo, err := ts.GetTodo(id)
	if err != nil {
		return nil, err
	}

	todo.Completed = !todo.Completed
	fmt.Println(todo)

	if err := ts.store.UpdateTodoCompleted(id, todo.Completed); err != nil {
		return nil, err
	}

	return todo, nil
}

func (ts *TodoService) DeleteTodo(id string) error {
	if err := ts.store.DeleteTodo(id); err != nil {
		return err
	}

	return nil
}
