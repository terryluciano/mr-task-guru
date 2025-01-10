package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terryluciano/mr-task-guru/db"
	"github.com/terryluciano/mr-task-guru/model"
	"github.com/terryluciano/mr-task-guru/utils"
)

func AddTaskHandler(c *gin.Context) {

	type newTask struct {
		Name           string  `json:"name"`
		Category       *int    `json:"category"`
		Minutes        *int    `json:"minutes"`
		Current_status *string `json:"current_status"`
	}

	var input model.Task

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Missing required information for new task",
		})
		return
	}

	err := db.InsertTask(&input)
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

	intID, err := utils.ConvertIdToInt(c, id)
	if err != nil {
		return
	}

	if err := db.RemoveTask(*intID); err != nil {
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

	var task model.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Request",
		})
		return
	}

	intID, err := utils.ConvertIdToInt(c, id)
	if err != nil {
		return
	}

	updateErr := db.UpdateTask(*intID, task)
	if updateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": updateErr.Error(),
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

	intID, err := utils.ConvertIdToInt(c, id)
	if err != nil {
		return
	}

	task, err := db.SelectTask(*intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, *task)
		return
	}

}

func GetAllTasksHandler(c *gin.Context) {
	tasks, err := db.SelectAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
	return
}
