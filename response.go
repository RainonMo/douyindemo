package douyindemo

import (
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

// 登录成功和注册成功
func Success(message string, userId uint64, token string, c *gin.Context) {
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: message},
		UserId:   userId,
		Token:    token,
	})
}