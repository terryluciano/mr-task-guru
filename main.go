package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlers "github.com/terryluciano/mr-task-guru/handlers"
	"github.com/terryluciano/mr-task-guru/pg"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	} else {
		log.Println("Env Variables Loaded...")
	}

	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	postgres_db := os.Getenv("POSTGRES_DB")
	postgres_host := os.Getenv("POSTGRES_HOST")
	postgres_port := os.Getenv("POSTGRES_PORT")

	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", postgres_user, postgres_password, postgres_host, postgres_port, postgres_db)

	pg.Connect(context.Background(), db_url)

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

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Println("ENV Port Not Found, Setting Default to 8080...")
		port = ":8080"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	router.Run(port)
}
