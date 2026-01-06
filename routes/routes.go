package routes

import "github.com/gin-gonic/gin"

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

	r.g.POST("/login", Login)

	r.g.GET("/profile/:username", GetProfile)

	r.g.POST("user/create", CreateUser)

	posts := r.g.Group("/posts")
	{
		posts.GET("/", GetPosts)
		posts.POST("/create", CreatePost)
		posts.POST("/:id/like", LikePost)
	}

}
