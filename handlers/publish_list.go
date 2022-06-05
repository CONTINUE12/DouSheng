package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

func PublishList(c *gin.Context) {
	var publishListParam request.PublishListParam

	if err := c.ShouldBind(&publishListParam); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed"})
	}

	if err := service.PublishList(publishListParam.UserId, publishListParam.Token); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed"})
		return
	}

	c.JSON(200, response.PublishList{
		Basic:     response.Basic{StatusCode: 0, StatusMsg: "success"},
		VideoList: repository.PublishList,
	})
}
