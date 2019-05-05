package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func ApiRouter() *gin.Engine {
	gin.DefaultWriter = colorable.NewColorableStdout()
	router := gin.Default()
	router.StaticFile("/robots.txt", "robots.txt")
	router.StaticFile("/favicon.ico", "favicon.ico")

	api := router.Group("/api/v1/")
	home := router.Group("/")

	// Categories
	// api.GET("/c", ListCategory) //LIST Categories
	// api.POST("/c", CreateCategory)      //CREATE Category
	// api.GET("/c/:id", RetrieveCategory) //RETRIEVE Category
	// api.PATCH("/c/:id", AlterCategory)  //ALTER Category
	// api.DELETE("/c/:id", PurgeCategory) //PURGE Category

	//Posts
	api.GET("/p", ListPosts)

	// To use auth routes you must have had a JWT token issued to you
	// and requests must have the following header:
	// -- Authorization: Bearer `token`
	// auth := api.Group("/auth/")
	// Use a JWT middleware authentication handler to only allow requests
	// that have been signed by our 'mysupersecretpassword' secret
	// auth.Use(jwt.Auth("mysupersecretpassword"))
	//auth.GET("/resetpassword", GetResetPassword)
	//auth.POST("/resetpassword", PostResetPassword)
	//auth.POST("/image", PostImage)

	home.GET("/ping", HomePage)

	return router
}

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
