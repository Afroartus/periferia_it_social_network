package main

import (
	database2 "periferia_it_social_network/database"
	models2 "periferia_it_social_network/models"
	"periferia_it_social_network/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database2.ConnectDB()

	database2.DB.AutoMigrate(
		&models2.User{},
		&models2.Post{},
	)
	database2.Seed()
	run()
}

func run() {
	g := gin.Default()
	r := routes.NewRoutes(g)
	r.SetupRoutes()
	g.Run(":8080")
}
