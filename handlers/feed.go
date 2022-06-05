package handlers

import (
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Feed(c *gin.Context) {
	latestStr := c.Query("latest_time")
	latestTime, _ := strconv.ParseInt(latestStr, 10, 64)
	token := c.Query("token")

	if err := service.GetLatestVideo(latestTime, token); err != nil {
		c.JSON(500, response.Basic{
			StatusCode: -1,
			StatusMsg:  "feed failed"})
		return
	}

	c.JSON(200, response.Feed{
		Basic:     response.Basic{StatusCode: 0, StatusMsg: "success"},
		VideoList: repository.LatestList,
		NextTime:  repository.NextTime,
	})
}
