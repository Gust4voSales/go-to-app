package controllers

import (
	"fmt"
	"net/http"

	"github.com/Gust4voSales/go-todo-app/internal/errors"
	"github.com/Gust4voSales/go-todo-app/internal/services"
	"github.com/Gust4voSales/go-todo-app/internal/types"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	tds *services.TodoService
}

func NewTodoController(tds *services.TodoService) *TodoController {
	return &TodoController{
		tds: tds,
	}
}

func (ctr *TodoController) GetTodosController(c *gin.Context) {
	todos, err := ctr.tds.ListTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errors.ErrorInternalServerError.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (ctr *TodoController) GetTodoController(c *gin.Context) {
	id := c.Param("id")

	// TODO add validation

	todo, err := ctr.tds.GetTodo(id)

	if err != nil {
		if err == errors.ErrorEntityNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("TODO with id %q not found", id),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errors.ErrorInternalServerError.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (ctr *TodoController) CreateTodoController(c *gin.Context) {
	var body types.CreateTodoBody

	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid content",
		})
		return
	}

	// TODO add validation

	todo, err := ctr.tds.CreateTodo(body.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func (ctr *TodoController) ToggleTodoCompletedController(c *gin.Context) {
	id := c.Param("id")

	// TODO add validation

	todo, err := ctr.tds.ToggleTodoCompleted(id)
	if err != nil {
		if err == errors.ErrorEntityNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("TODO with id %q not found", id),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": errors.ErrorInternalServerError.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (ctr *TodoController) DeleteTodoController(c *gin.Context) {
	id := c.Param("id")

	// TODO add validation

	if err := ctr.tds.DeleteTodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errors.ErrorInternalServerError.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
