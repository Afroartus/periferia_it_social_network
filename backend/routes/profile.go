package routes

import (
	"log"
	"net/http"
	"periferia_it_social_network/database"
	"periferia_it_social_network/models"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	username := c.GetString("username")

	log.Println(username)
	
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No user found",
			"code":    http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user.Username,
	})
}

func GetProfile(c *gin.Context) {
	var user ProfileResponse

	userName := c.Param("username")

	err := database.DB.Model(&models.User{}).
		Where("username = ?", userName).Find(&user).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error: user not found",
			"code":    http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
