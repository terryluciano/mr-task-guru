package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/terryluciano/mr-task-guru/utils"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func AddTaskHandler(c *gin.Context) {

	var newTask utils.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Missing required information for new task",
		})
		return
	}

	err := utils.AddTask(newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed to add new task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func GetAllTasksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, utils.GetAllTask())
}
