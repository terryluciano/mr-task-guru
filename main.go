package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlers "github.com/terryluciano/mr-task-guru/handlers"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("Env Variables Loaded")
	}

	router := gin.Default()

	// Tasks router
	tasks := router.Group("/task")
	{
		tasks.POST("/add", handlers.AddTaskHandler)
		tasks.GET("/all", handlers.GetAllTasksHandler)
		tasks.GET("/:id", handlers.GetTaskHandler)
		tasks.DELETE("/remove/:id", handlers.RemoveTaskHandler)
		tasks.PUT("/update/:id", handlers.UpdateTaskHandler)
	}

	// Category router
	category := router.Group("/category")
	{
		category.POST("/add/:category", handlers.AddCategoryHandler)
		category.DELETE("/remove/:id", handlers.RemoveCategoryHandler)
		category.GET("/:id", handlers.GetCategoryHandler)
		category.GET("/all", handlers.GetAllCategoriesHandler)
	}

	router.Run(":4000")
}
