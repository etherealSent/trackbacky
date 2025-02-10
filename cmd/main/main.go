package main

import (
	"awesomeProject/web-service-gin/cmd/main/app"
	_ "awesomeProject/web-service-gin/cmd/main/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"github.com/joho/godotenv"
    "log"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	rt := gin.Default()

	rt.POST("/getsong", app.GetSong)
	rt.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
	}))

	rt.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://176.114.67.228:8089/swagger/doc.json")))

    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = "8088"
    }

	rt.Run("0.0.0.0" + ":" + port)

}
