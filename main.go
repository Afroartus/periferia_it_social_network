package main

import (
	"periferia_it_social_network/database"

	"github.com/gin-gonic/gin"
)

func main() {
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
