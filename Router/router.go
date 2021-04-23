package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/hewenyu/gin-simple/Controllers/hello_demo"
)

// GetUser
func InitAuthRouter(Router *gin.RouterGroup) {
	AuthGroup := Router.Group("test")
	{
		AuthGroup.GET("", hello_demo.GetHellow) // 通过令牌获取token
	}
}
