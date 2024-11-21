package services

import (
	"fmt"
	"strconv"

	"github.com/Gust4voSales/go-todo-app/internal/types"
)

type TodoService struct {
	todoList []types.Todo
}

func New(tl []types.Todo) *TodoService {
	return &TodoService{
		todoList: tl,
	}
}

func (ts *TodoService) ListTodos() []types.Todo {
	return ts.todoList
}

func (ts *TodoService) GetTodo(id string) (*types.Todo, error) {
	for i, v := range ts.todoList {
		if v.ID == id {
			return &ts.todoList[i], nil
		}
	}
	return nil, fmt.Errorf("todo with ID: %q not found", id)
}

func (ts *TodoService) CreateTodo(content string) types.Todo {
	ID := len(ts.todoList) + 1

	todo := types.Todo{
		ID:        strconv.Itoa(ID),
		Content:   content,
		Completed: false,
	}

	ts.todoList = append(ts.todoList, todo)

	return todo
}

func (ts *TodoService) ToggleTodoCompleted(id string) (*types.Todo, error) {
	todo, err := ts.GetTodo(id)
	if err != nil {
		return nil, err
	}

	todo.Completed = !todo.Completed

	return todo, nil
}

func (ts *TodoService) DeleteTodo(id string) error {
	for i, v := range ts.todoList {
		if v.ID == id {
			ts.todoList = append(ts.todoList[:i], ts.todoList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo with ID: %q not found", id)
}
