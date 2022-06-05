package handlers

import (
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	if err := service.GetCommentList(videoId, token); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to get comment list"})
		return
	}

	c.JSON(200, response.CommentList{
		Basic:       response.Basic{StatusCode: 0, StatusMsg: "success"},
		CommentList: repository.CommentList,
	})

}
