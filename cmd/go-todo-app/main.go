package main

import (
	"log"

	"github.com/Gust4voSales/go-todo-app/internal/controllers"
	"github.com/Gust4voSales/go-todo-app/internal/db"
	"github.com/Gust4voSales/go-todo-app/internal/services"
	"github.com/Gust4voSales/go-todo-app/internal/stores"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	g := gin.Default()

	db, err := db.NewConnection()
	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}

	todoStore := stores.NewTodoStore(db)

	todoService := services.NewTodoService(todoStore)

	todoController := controllers.NewTodoController(todoService)

	// ROUTES
	g.GET("/todos", todoController.GetTodosController)
	g.GET("/todos/:id", todoController.GetTodoController)
	g.POST("/todos", todoController.CreateTodoController)
	g.PATCH("/todos/:id/toggle-completed", todoController.ToggleTodoCompletedController)
	g.DELETE("/todos/:id", todoController.DeleteTodoController)

	// RUN SERVER
	g.Run("localhost:3000")
}
