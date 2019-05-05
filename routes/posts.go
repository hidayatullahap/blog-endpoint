package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hidayatullahap/blog-endpoint/models"
)

func ListPosts(c *gin.Context) {
	posts := models.GetPosts()
	c.JSON(200, posts)
}
