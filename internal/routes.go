package internal

import (
	"fmt"
	"net/http"

	"github.com/Gust4voSales/go-todo-app/internal/services"
	"github.com/Gust4voSales/go-todo-app/internal/types"
	"github.com/gin-gonic/gin"
)

var (
	todoList    []types.Todo
	todoService *services.TodoService
)

func SetupRoutes(g *gin.Engine) {
	todoList = []types.Todo{}
	todoService = services.New(todoList)

	g.GET("/todos", getTodosController)
	g.GET("/todos/:id", getTodoController)
	g.POST("/todos", createTodoController)
	g.PATCH("/todos/:id/toggle-completed", toggleTodoCompletedController)
	g.DELETE("/todos/:id", deleteTodoController)
}

func getTodosController(c *gin.Context) {
	todos := todoService.ListTodos()

	c.JSON(http.StatusOK, todos)
}

func getTodoController(c *gin.Context) {
	id := c.Param("id")

	// TODO add validation

	todo, err := todoService.GetTodo(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func createTodoController(c *gin.Context) {
	var body types.CreateTodoBody

	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid content",
		})
		return
	}

	// TODO add validation

	todo := todoService.CreateTodo(body.Content)

	c.JSON(http.StatusCreated, todo)
}

func toggleTodoCompletedController(c *gin.Context) {
	id := c.Param("id")

	// TODO add validation

	todo, err := todoService.ToggleTodoCompleted(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func deleteTodoController(c *gin.Context) {
	id := c.Param("id")

	// TODO add validation

	if err := todoService.DeleteTodo(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
