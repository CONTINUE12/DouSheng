package handlers

import (
	"douyin/http/response"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Favorite(c *gin.Context) {

	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)

	if err := service.AddFavorite(userId, videoId, actionType, token); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to favorite"})
		return
	}

	c.JSON(200, response.Basic{StatusCode: 0, StatusMsg: "success"})
}
