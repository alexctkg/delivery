package main

import (
	"delivery/controllers"
	"delivery/docs"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	gotenv.Load()
	gin.SetMode(os.Getenv("GIN_MODE"))

	// Swagger configuration
	docs.SwaggerInfo.Title = "Delivery"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http, https"}

	router := SetupRouter()
	router.Run()

}

//SetupRouter configuration
func SetupRouter() *gin.Engine {
	gotenv.Load()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "GET", "DELETE", "POST"},
		AllowHeaders:    []string{"Content-type", "Authorization"},
		ExposeHeaders:   []string{"Content-Length", "Content-type"},
		MaxAge:          36000,
	}))

	router.GET("/recipes/", controllers.IndexRecipe)

	swagger := router.Group("/docs", gin.BasicAuth(gin.Accounts{
		"delivery": "1234",
	}))
	swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // api documentation HOST/docs/index.html

	return router
}
