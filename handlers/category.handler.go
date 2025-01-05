package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/terryluciano/mr-task-guru/utils"
)

func AddCategoryHandler(c *gin.Context) {
	newCategory := c.Param("category")

	if newCategory, err := utils.AddCategory(newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Successfully Added a New Category",
			"data": newCategory,
		})
		return
	}
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

func GetCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	if category, err := utils.GetCategory(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, category)
		return
	}
}

func GetAllCategoriesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, utils.GetAllCategories())
}
