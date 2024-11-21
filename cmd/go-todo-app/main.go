package main

import (
	"github.com/Gust4voSales/go-todo-app/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	internal.SetupRoutes(g)

	g.Run(":3000")
}
