package types

type Todo struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

type CreateTodoBody struct {
	Content string `json:"content"`
}
