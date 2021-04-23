package hello_demo

import (
	"github.com/gin-gonic/gin"
	"github.com/hewenyu/gin-simple/Models/response"
	"github.com/hewenyu/gin-simple/Services/demo"
)

// @Tags Test
// @Summary 获取当前时间
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response"
// @Router /api/test [GET]
func GetHellow(c *gin.Context) {

	now_data := demo.HelloWorld()

	response.OkWithDetailed(now_data, "获取成功", c)

}
