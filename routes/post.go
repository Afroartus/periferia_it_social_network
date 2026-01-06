package routes

import (
	"net/http"
	"periferia_it_social_network/database"
	"periferia_it_social_network/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Order("created_at desc").Find(&posts)
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func CreatePost(c *gin.Context) {
	user := c.GetUint("user_id")

	var request CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "message required",
			"code":  http.StatusBadRequest,
		})
		return
	}

	post := models.Post{
		Message: request.Message,
		UserID:  user,
	}

	database.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "post created",
		"code":    http.StatusCreated,
	})
}

func LikePost(c *gin.Context) {
	id := c.Param("id")

	err := database.DB.Model(&models.Post{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes + 1")).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "post not found",
			"code":  http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post liked",
	})
}
