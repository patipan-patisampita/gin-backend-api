package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/patipan-patisampita/gin-backend-api/routes"
)

func main() {
	router := SetupRouter()
	router.Run(":" + os.Getenv("GO_PORT")) //Listen and Server on 0.0.0.0:8080
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiV1 := router.Group("api/v1")

	routes.InitHomeRoutes(apiV1)
	routes.InitUserRoutes(apiV1)
	return router
}
