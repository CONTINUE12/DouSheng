package handlers

import (
	"douyin/http/response"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Action 添加关注或取消关注
func Action(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)

	fmt.Printf("%v", userId)
	if err := service.Relation(userId, toUserId, actionType, token); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed of post action"})
		return
	}

	c.JSON(200, response.Basic{StatusCode: 0, StatusMsg: "success"})

}
