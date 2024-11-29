package stores

import (
	"database/sql"
	"fmt"

	"github.com/Gust4voSales/go-todo-app/internal/errors"
	"github.com/Gust4voSales/go-todo-app/internal/types"
)

type TodoStore struct {
	db *sql.DB
}

func NewTodoStore(db *sql.DB) TodoStoreInterface {
	return &TodoStore{
		db: db,
	}
}

func (s *TodoStore) ListTodos() ([]types.Todo, error) {
	rows, err := s.db.Query("SELECT * FROM todos")

	if err != nil {
		fmt.Println(err.Error())
		return []types.Todo{}, err
	}

	todos := []types.Todo{}
	var todoTemp types.Todo

	for rows.Next() {
		err = rows.Scan(
			&todoTemp.ID,
			&todoTemp.Content,
			&todoTemp.Completed,
		)

		if err != nil {
			return []types.Todo{}, err
		}

		todos = append(todos, todoTemp)
	}

	rows.Close()

	return todos, nil
}

func (s *TodoStore) GetTodo(id string) (*types.Todo, error) {
	row := s.db.QueryRow("SELECT * FROM todos WHERE id = $1", id)

	var todo types.Todo

	if err := row.Scan(&todo.ID, &todo.Content, &todo.Completed); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorEntityNotFound
		}
		return nil, err
	}

	return &todo, nil
}

func (s *TodoStore) CreateTodo(todo types.Todo) error {
	if _, err := s.db.Exec("INSERT INTO todos VALUES ($1, $2, $3)", todo.ID, todo.Content, todo.Completed); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (s *TodoStore) UpdateTodoCompleted(id string, status bool) error {
	if _, err := s.db.Exec("UPDATE todos SET completed = $1 WHERE id = $2", status, id); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *TodoStore) DeleteTodo(id string) error {
	if _, err := s.db.Exec("DELETE FROM todos WHERE id = $1", id); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
