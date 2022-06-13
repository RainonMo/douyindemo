package main

import (
	"douoooyindemo/common"
	"douoooyindemo/controller"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)
var Db *gorm.DB
func main() {
	common.Mysql()
	
	engine := gin.Default()

	engine.Static("/static", "./public")

	// 后台管理员前端接口
	web := engine.Group("/douyin")
	

	{
		//测试
		web.GET("/test",func(c *gin.Context) {
			c.JSON(200,gin.H{
				"msg":"hello",
			})
		})
		// 用户登录API
		 web.POST("/user/login/", controller.Login)
		 web.GET("/user/", controller.UserInfo)
		 web.POST("/user/register/", controller.Register)


	}

	// 启动
	engine.Run()
	
	

}