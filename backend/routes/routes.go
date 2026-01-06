package routes

import (
	"periferia_it_social_network/routes/middleware"
	"time"

	"github.com/gin-contrib/cors"
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

	r.g.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Authorization", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length"},
		AllowAllOrigins: true,
		MaxAge:          12 * time.Hour,
	}))

	r.g.GET("/ping", Ping)
	r.g.POST("/login", Login)
	r.g.POST("user/create", CreateUser)

	auth := r.g.Group("/")
	auth.Use(middleware.Auth())
	{
		auth.GET("/profile/:username", GetProfile)
		auth.GET("/me", GetMe)

		posts := auth.Group("/posts")
		{
			posts.GET("", GetPosts)
			posts.POST("", CreatePost)
			posts.POST("/:id/like", LikePost)
		}

	}

}
