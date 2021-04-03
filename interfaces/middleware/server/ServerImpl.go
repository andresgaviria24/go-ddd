package server

import (
	"log"
	"os"
	"ws_restaurant/application"
	"ws_restaurant/docs"
	"ws_restaurant/interfaces/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerImpl struct {
	router *gin.Engine
}

func InitServer() Server {
	serverImpl := &ServerImpl{}
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	swaggerDocs()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	application.InitRestaurantController(router)
	serverImpl.router = router
	return serverImpl
}

func swaggerDocs() {
	docs.SwaggerInfo.Title = os.Getenv("SWAGGER_TITLE")
	docs.SwaggerInfo.Description = os.Getenv("SWAGGER_DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("SWAGGER_VERSION")
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASEPATH")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func (api *ServerImpl) RunServer() {
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = os.Getenv("LOCAL_PORT") //localhost
	}
	log.Fatal(api.router.Run(":" + appPort))
}
