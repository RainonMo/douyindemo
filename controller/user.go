package controller

import (
	douyindemo "douoooyindemo"
	"douoooyindemo/entity"
	"douoooyindemo/service"

	"github.com/gin-gonic/gin"
)

var userService service.WebUserService

//登录
func Login(c *gin.Context){
	var param entity.AdminWebUserLoginVO
	param.Username = c.Query("username")
	param.Password = c.Query("password")

	uid := userService.Login(param)

	//找到用户
	if uid>0{
		//todo token
		douyindemo.Success("登录成功",uid,token,c)
	}
}