package routes

import (
	"periferia_it_social_network/routes/middleware"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	g *gin.Engine
}

func NewRoutes(engine *gin.Engine) *Routes {
	return &Routes{
		g: engine,
	}
}

func (r *Routes) SetupRoutes() {
	r.g.GET("/ping", Ping)

	r.g.POST("user/create", CreateUser)
	r.g.POST("/login", Login)

	auth := r.g.Group("/")
	auth.Use(middleware.Auth())
	{
		auth.GET("/profile/:username", GetProfile)

		posts := auth.Group("/posts")
		{
			posts.GET("/", GetPosts)
			posts.POST("/create", CreatePost)
			posts.POST("/:id/like", LikePost)
		}

	}

}
