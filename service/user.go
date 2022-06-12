package service

import (
	"douoooyindemo/common"
	"douoooyindemo/dao"
	"douoooyindemo/entity"

	"github.com/golang-module/carbon/v2"
)

//service接口WebUserService
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

//根据token获取用户
func (u *WebUserService) GetUserByToken(token string) (entity.User, error) {
	var user entity.User

	uid, err := common.GetUidByToken(token)
	if err != nil {
		return user, err

	}

	user = dao.UserMgr.GetInfoById(uid)

	return user, nil
}
