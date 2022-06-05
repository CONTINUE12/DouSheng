package handlers

import (
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Comment(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	commentText := c.Query("comment_text")
	deleteId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

	if actionType == 1 {
		if err := service.AddComment(userId, videoId, token, commentText); err != nil {
			c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to add comment"})
		} else {
			c.JSON(200, response.CommentAction{
				Basic:   response.Basic{StatusCode: 0, StatusMsg: "success to add comment"},
				Comment: repository.ActionComment,
			})
		}
	}
	if actionType == 2 {
		if err := service.DeleteComment(deleteId, token); err != nil {
			c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to delete comment"})
		} else {
			c.JSON(200, response.CommentAction{
				Basic: response.Basic{StatusCode: 0, StatusMsg: "success to delete comment"},
			})
		}
	}
}
