package main

import (
	"github.com/Gust4voSales/go-todo-app/internal/controllers"
	"github.com/Gust4voSales/go-todo-app/internal/services"
	"github.com/Gust4voSales/go-todo-app/internal/types"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	todoList := []types.Todo{}
	todoService := services.NewTodoService(todoList)
	todoController := controllers.NewTodoController(todoService)

	g.GET("/todos", todoController.GetTodosController)
	g.GET("/todos/:id", todoController.GetTodoController)
	g.POST("/todos", todoController.CreateTodoController)
	g.PATCH("/todos/:id/toggle-completed", todoController.ToggleTodoCompletedController)
	g.DELETE("/todos/:id", todoController.DeleteTodoController)

	g.Run("localhost:3000")
}
