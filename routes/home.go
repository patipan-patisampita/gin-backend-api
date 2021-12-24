package routes

import "github.com/gin-gonic/gin"

func InitHomeRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/")

	//http://localhost:xx/api/v1/
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API VESRION": "1.0.0",
		})
	})
}
