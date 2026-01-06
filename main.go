package main

import (
	"periferia_it_social_network/database"
	"periferia_it_social_network/models"
	"periferia_it_social_network/routes"

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
	g := gin.Default()
	r := routes.NewRoutes(g)
	r.SetupRoutes()
	g.Run(":8080")
}
