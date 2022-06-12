package controller

import (
	"douoooyindemo/common"
	"douoooyindemo/entity"
	"douoooyindemo/service"

	"github.com/gin-gonic/gin"
)

//controller的userService注入services的接口
var userService service.WebUserService

//登录
func Login(c *gin.Context){
	var param entity.AdminWebUserLoginVO
	param.Username = c.Query("username")
	param.Password = c.Query("password")

	uid := userService.Login(param)

	//找到用户
	if uid>0{
		//token
		token,err := common.GenerateToke(param.Username,uid)
		if err !=nil{
			common.Failed("生成token失败",c)
			return
		}
		common.Success("登录成功",uid,token,c)
		return
		
	}

	common.Failed("未找到用户",c)
}

//注册
func Register(c *gin.Context){
	var param entity.AdminWebUserLoginVO

	param.Username = c.Query("username")
	param.Password = c.Query("password")

	flag,user := userService.Register(param)

	if flag{
		//token
		token,err := common.GenerateToke(user.Name,user.Id)
		if err !=nil{
			common.Failed("生成token失败",c)
			return
		}
		common.Success("注册成功",user.Id,token,c)

	}else{
		common.Failed("用户已存在",c)
	}
}

//用户信息
func UserInfo(c *gin.Context){
	//token
	token := c.Query("token")
	user,err := userService.GetUserByToken(token)
	if err != nil{
		common.Failed(err.Error(),c)
		return
	}
	if user.Id >0{
		common.UserInfo(user,c)
		return
	}
	common.Failed("用户不存在",c)

}