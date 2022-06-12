package controller

import (
	"time"
	"net/http"
	"douoooyindemo/common"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context){
	c.JSON(http.StatusOK,common.FeedResponse{
		Response:  common.Response{StatusCode: 0},
		//VideoList: global.VideoList,
		NextTime:  time.Now().Unix(),
	})
}