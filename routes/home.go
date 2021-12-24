package routes

import "github.com/gin-gonic/gin"

func InitHomeRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")

	//http://localhost:xx/api/v1/users
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "users",
		})
	})
}
