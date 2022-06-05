package handlers

import (
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FavoriteList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")

	if err := service.GetFavoriteList(userId, token); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to get favorite list"})
		return
	}

	c.JSON(200, response.FavoriteList{
		Basic:     response.Basic{StatusCode: 0, StatusMsg: "success"},
		VideoList: repository.FavoriteList,
	})
}
