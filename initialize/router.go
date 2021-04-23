package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hewenyu/gin-simple/Middlewares"
	"github.com/hewenyu/gin-simple/Router"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由
func Routers() *gin.Engine {
	var DefultRouter = gin.Default()

	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	// Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	DefultRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 不需要验证的路由
	PublicGroup := DefultRouter.Group("/api")
	{
		Router.InitAuthRouter(PublicGroup) // 测试接口

	}
	// 需要验证的路由
	PrivateGroup := DefultRouter.Group("/api")
	PrivateGroup.Use(Middlewares.JWTAuth())
	// {
	// 	Router.InitAuthRouter(PrivateGroup) // 测试接口
	// }
	return DefultRouter
}
