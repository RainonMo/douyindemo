package dao

import (
	douyindemo "douoooyindemo"
	"douoooyindemo/entity"
)

type UserManager interface {
	Register(user *entity.User)
	Login(name, password string) entity.User
	GetInfoById(uid uint64) entity.User
	GetUserByName(name string) entity.User
}

type manager struct {
}

var UserMgr UserManager = &manager{}

//登录
func (mgr *manager) Login(name, password string) entity.User {
	var user entity.User

	douyindemo.Db.Where("name = ? and password = ?", name, password).First(&user)
	return user
}

//注册
func (mgr *manager) Register(user *entity.User) {
	douyindemo.Db.Create(user)
}

//根据id查询
func (mgr *manager) GetInfoById(uid uint64) entity.User {
	var user entity.User

	douyindemo.Db.Where("id = ?", uid).First(&user)
	return user
}

//根据name查询
func (mgr *manager) GetUserByName(name string) entity.User {
	var user entity.User

	douyindemo.Db.Where("name = ?", name).First(&user)
	return user
}