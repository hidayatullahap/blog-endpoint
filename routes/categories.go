package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create categories",
	})
}

func RetrieveCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "retrieve categories",
	})
}

func AlterCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "alter categories",
	})
}

func PurgeCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "purge categories",
	})
}

func ListCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list categories",
	})
}
