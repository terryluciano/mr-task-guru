package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terryluciano/mr-task-guru/db"
	"github.com/terryluciano/mr-task-guru/utils"
)

func AddCategoryHandler(c *gin.Context) {
	category := c.Param("category")

	if newCategory, err := db.InsertCategory(category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Successfully Added a New Category",
			"data": *newCategory,
		})
		return
	}
}

func RemoveCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	inputID, err := utils.ConvertIdToInt(c, id)
	if err != nil {
		return
	}

	if err := db.RemoveCategory(*inputID); err != nil {
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

	inputID, err := utils.ConvertIdToInt(c, id)
	if err != nil {
		return
	}

	if category, err := db.SelectCategory(*inputID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, *category)
		return
	}
}

func GetAllCategoriesHandler(c *gin.Context) {
	categories, err := db.SelectAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
	return
}
