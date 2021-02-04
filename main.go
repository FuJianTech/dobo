package main

import (
	// "fmt"

	"dobo/config"
	"dobo/database"
	"dobo/middleware"
	"dobo/routers"
	"dobo/utils"
)

var configure = config.ConfigMain()

// @title tool工具接口
// @version 1.0
// @description 接口信息
// @termsOfService http://swagger.io/terms/
// @contact.name HZhenyong
// @contact.url http://www.swagger.io/support
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @contact.email hzy01hzy@gmail.com
// @host localhost:8083
// @BasePath /api/v1
func main() {
	db := database.InitDB()
	defer db.Close()

	middleware.MySigningKey = []byte(configure.Server.ServerJWT.Key)
	middleware.JWTExpireTime = configure.Server.ServerJWT.Expire


	router := routers.SetupRouter()
	err := router.Run(":" + configure.Server.ServerPort)
	if err != nil {
		utils.Logger().Error(err)
		//fmt.Printf("startup service failed, err:%v\n\n", err)
	}
	utils.Logger().Info(" Listening and serving HTTP on :"+configure.Server.ServerPort)
}

