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

	// 开启跨域
	//engine.Use(middleware.Cors())

	// 静态资源请求映射
	// engine.Static("/video", global.Config.Upload.SavePath)
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

		web.GET("/feed/", controller.Feed)
		//web.POST("/publish/action/", controller.Publish)

		// 开启JWT认证，以下接口需要认证成功才能访问
		//web.Use(middleware.JwtAuth())

	}

	// 启动
	engine.Run()
	
	

}