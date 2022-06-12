package common

import (
	"douoooyindemo/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	// Data    interface{} `json:"data"`
}

// 用户的登录成功
type UserLoginResponse struct {
	Response
	UserId uint64 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// 返回用户信息
type UserResponse struct {
	Response
	User entity.User `json:"user"`
}

// FeedResponse
type FeedResponse struct {
	Response
	VideoList []entity.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// 登录成功和注册成功
func Success(message string, userId uint64, token string, c *gin.Context) {
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: message},
		UserId:   userId,
		Token:    token,
	})
}

// Failed 请求失败返回
func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 1, StatusMsg: message},
	})
}

// 返回用户信息
func UserInfo(user entity.User, c *gin.Context) {
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0,StatusMsg: "null"},
		User:     user,
	})
}