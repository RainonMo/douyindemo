package service

import (
	"douoooyindemo/dao"
	"douoooyindemo/entity"

	"github.com/golang-module/carbon/v2"
)

type WebUserService struct{}

//登录
func (u *WebUserService) Login(param entity.AdminWebUserLoginVO) uint64{
	user := dao.UserMgr.Login(param.Username,param.Password)
	return user.Id
}

//注册
func (u *WebUserService) Register(param entity.AdminWebUserLoginVO) (bool,*entity.User){
	//查询name
	user := dao.UserMgr.GetUserByName(param.Username)

	if user.Id > 0{
		return false,nil
	}else {
		//创建新用户
		user.Name = param.Username
		user.Password = param.Password
		user.CreateAt = carbon.Now().ToDateTimeString()
		user.UpdateAt = carbon.Now().ToDateTimeString()
		dao.UserMgr.Register(&user)

		user = dao.UserMgr.GetUserByName(param.Username)

		return true,&user
	}
}

//token