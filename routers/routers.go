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
		// Register - no JWT required
		//v1.POST("/register",middleware.JWT(), controller.CreateUserAuth)

		// Login - app issues JWT
		//v1.POST("/login", controller.Login)

		//// User
		//rUsers := v1.Group("users")
		//rUsers.GET("", controller.GetUsers)    // Non-protected
		//rUsers.GET("/:id", controller.GetUser) // Non-protected
		//rUsers.Use(middleware.JWT())
		//rUsers.POST("", controller.CreateUser)      // Protected
		//rUsers.PUT("", controller.UpdateUser)       // Protected
		//rUsers.PUT("/hobbies", controller.AddHobby) // Protected
		//
		//// Post
		//rPosts := v1.Group("posts")
		//rPosts.GET("", controller.GetPosts)    // Non-protected
		//rPosts.GET("/:id", controller.GetPost) // Non-protected
		//rPosts.Use(middleware.JWT())
		//rPosts.POST("", controller.CreatePost)       // Protected
		//rPosts.PUT("/:id", controller.UpdatePost)    // Protected
		//rPosts.DELETE("/:id", controller.DeletePost) // Protected
		//
		//// Hobby
		//rHobbies := v1.Group("hobbies")
		//rHobbies.GET("", controller.GetHobbies) // Non-protected
	}

	return router
}



