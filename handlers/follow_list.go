package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

func FollowList(c *gin.Context) {
	var followParam request.RelationParam

	if err := c.ShouldBind(&followParam); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "bind param failed"})
	}

	if err := service.FollowList(followParam.UserId, followParam.Token); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to get follow"})
		return
	}

	c.JSON(200, response.FollowList{
		Basic:    response.Basic{StatusCode: 0, StatusMsg: "success"},
		UserList: repository.FollowUsers,
	})
}
