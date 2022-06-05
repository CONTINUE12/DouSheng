package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

func FollowerList(c *gin.Context) {
	var followerParam request.RelationParam

	if err := c.ShouldBind(&followerParam); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "bind param failed"})
	}

	if err := service.FollowerList(followerParam.UserId, followerParam.Token); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to get follower list"})
		return
	}

	c.JSON(200, response.FollowerList{
		Basic:    response.Basic{StatusCode: 0, StatusMsg: "success"},
		UserList: repository.FollowerUsers,
	})
}
