package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUp struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func LoginHandler(c *gin.Context) {
	var loginData Login

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Missing required information to login",
		})
		return
	}

}
