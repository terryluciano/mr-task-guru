package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/terryluciano/mr-task-guru/utils"
)

func AddCategoryHandler(c *gin.Context) {
	newCategory := c.Param("category")

	err := utils.AddCategory(newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Successfully Added a New Category",
	})
	return
}

func RemoveCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	if err := utils.RemoveCategory(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Successfully Removed Category",
	})
	return
}
