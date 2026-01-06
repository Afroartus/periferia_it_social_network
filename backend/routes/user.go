package routes

import (
	"net/http"
	"periferia_it_social_network/database"
	"periferia_it_social_network/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var userRequest UserRequest
	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user required",
			"code":  http.StatusBadRequest,
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userRequest.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
			"code":  http.StatusInternalServerError,
		})
	}

	birth, err := time.Parse("2006-01-02", userRequest.BirthDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "birth date required",
			"code":  http.StatusBadRequest,
		})

		return
	}

	user := models.User{
		Name:      userRequest.Name,
		LastName:  userRequest.Lastname,
		Username:  userRequest.Username,
		Email:     userRequest.Email,
		Password:  string(hashPassword),
		BirthDate: birth,
	}

	database.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{
		"user": "user created",
		"code": http.StatusCreated,
	})
}
