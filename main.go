package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/terryluciano/mr-task-guru/handlers"
)

func main() {
	router := gin.Default()

	// Tasks
	tasks := router.Group("/task")
	{
		tasks.POST("/add", handlers.AddTaskHandler)
		tasks.GET("/all", handlers.GetAllTasksHandler)
		tasks.GET("/:id", handlers.GetTaskHandler)
		tasks.DELETE("/remove/:id", handlers.RemoveTaskHandler)
		tasks.PUT("/update/:id", handlers.UpdateTaskHandler)
	}

	// Categories router
	category := router.Group("/category")
	{
		category.POST("/add/:category", handlers.AddCategoryHandler)
		category.DELETE("/remove/:id", handlers.RemoveCategoryHandler)
	}

	router.Run(":4000")
}
