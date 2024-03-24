package main

import (
	"github.com/gin-gonic/gin"

	_ "mygram-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"mygram-api/routes"
)


func main() {

	router := gin.Default()

	// Mount Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Set up routes
	routes.UserRoute(router)
	routes.PhotoRoute(router)
	routes.CommentRoute(router)
	routes.SocialMediaRoute(router)
	
	router.Run(":8080")

}
