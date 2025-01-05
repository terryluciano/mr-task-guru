package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/terryluciano/mr-task-guru/utils"
)

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
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Successfully Added Task",
	})
	return
}

// task id is in the param
func RemoveTaskHandler(c *gin.Context) {

	id := c.Param("id")
	if err := utils.RemoveTask(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Successfully Removed Task",
		})
		return
	}

}

// will take either status, name, category, and/or minutes, but needs to pass an id
func UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")

	type UpdateTaskRequest struct {
		Name     *string       `json:"name"`
		Minutes  *int          `json:"minutes"`
		Category *string       `json:"category"`
		Status   *utils.Status `json:"status"`
	}

	var task UpdateTaskRequest
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Request",
		})
		return
	}

	err := utils.UpdateTask(id, task.Name, task.Minutes, task.Category, task.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Successfully Updated Task",
	})
	return
}

func GetTaskHandler(c *gin.Context) {
	id := c.Param("id")

	task, err := utils.GetTask(id)
	if err != nil {
		c.JSON(http.StatusFound, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, task)
		return
	}

}

func GetAllTasksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, utils.GetAllTask())
	return
}
