package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/terryluciano/mr-task-guru/handlers"
)

func main() {
	router := gin.Default()
	router.POST("/task/add", handlers.AddTaskHandler)
	router.GET("/task/all", handlers.GetAllTasksHandler)
	router.GET("/ping", handlers.PingHandler)
	router.Run(":4000")
}
