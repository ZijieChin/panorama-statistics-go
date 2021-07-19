package router

import (
	v1 "Panorama-Statistics/controller/v1"
	validator "Panorama-Statistics/middleware/validator"

	_ "Panorama-Statistics/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))
	r.Use(validator.Validate())
	GroupV1 := r.Group("v1")
	{
		GroupV1.POST("/stat", v1.InsertInfo)
		// params: user, page, start, end
		GroupV1.GET("/stat", v1.GetInfo)
		GroupV1.GET("/file", v1.GetFile)
		GroupV1.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	}
	return r
}
