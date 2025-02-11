package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/terryluciano/mr-task-guru/server/services"
	"github.com/terryluciano/mr-task-guru/server/utils"
)

func AddCategoryHandler(c *gin.Context) {
	type newCategory struct {
		Name  string  `json:"name"`
		Color *string `json:"color"`
	}

	var input newCategory

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Missing required information for new category",
		})
		return
	}

	if newCategory, err := db.InsertCategory(input.Name, input.Color); err != nil {
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
