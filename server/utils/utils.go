package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TitleizeString(str string) string {
	caser := cases.Title(language.English)
	return caser.String(str)
}

func GenerateRandomID() (string, error) {
	currentTime := time.Now().UnixNano()
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", fmt.Errorf("Could not generate id: %w", err)
	}

	newID := strconv.Itoa(int(currentTime)) + randomNumber.String()

	return newID, nil

}

func ConvertIdToInt(c *gin.Context, id string) (*int, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "ID must be a number",
		})
		return nil, fmt.Errorf("ID must be a number")
	}
	return &intID, nil
}
