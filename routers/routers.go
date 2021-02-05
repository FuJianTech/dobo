package routers

import (

	"dobo/config"
	"dobo/controller"
	_ "dobo/docs"
	"dobo/middleware"
	"dobo/utils"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)


var configure = config.ConfigMain()
// SetupRouter ...
func SetupRouter() *gin.Engine {
	if configure.Server.ServerEnv == "production" {
		gin.SetMode(gin.ReleaseMode) // Omit this line to enable debug mode
	}


	router := gin.Default()
	router.Use(utils.LoggerToFile())
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.CORS())

	// API:v1.0

	v1 := router.Group("/api/v1")
	{

		rAuth := v1.Group("")
		rAuth.POST("/login", controller.Login)
		rAuth.Use(middleware.JWT())
		rAuth.POST("/register",controller.CreateUserAuth)


		rDocker := v1.Group("/docker")
		rDocker.Use(middleware.JWT())
		rDocker.GET("/dockerInfo",controller.DockerInfo)
		rDocker.GET("/getVersion", controller.GerVersion)
		rDocker.GET("/ping", controller.Ping)
		rDocker.GET("/diskUsage", controller.DiskUsage)

		rImage := v1.Group("/images")
		rImage.Use(middleware.JWT())
		rImage.GET("/getImageList", controller.GetImgesList)
		rImage.GET("/getImageInfo/:imageId", controller.GetImageInfo)
		rImage.GET("/deleteImage/:imageId", controller.DeleteImage)
		rImage.GET("/reTagImage", controller.ReTagImage)


	}

	return router
}



