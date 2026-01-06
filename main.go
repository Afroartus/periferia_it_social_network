package main

import (
	"periferia_it_social_network/database"
	"periferia_it_social_network/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	database.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
	)
	database.Seed()
	run()
}

func run() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
