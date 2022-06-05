package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var registerParam request.UserParam

	//绑定参数
	if err := c.ShouldBind(&registerParam); err != nil {
		print("bind param err")
		c.JSON(401, response.Register{
			Basic: response.Basic{
				StatusCode: -1,
				StatusMsg:  "bind param err",
			},
			UserId: 1,
			Token:  "zxy",
		})
		return
	}

	//校验参数
	if len(registerParam.PassWord) == 0 || len(registerParam.UserName) == 0 {
		print("param length err")
		c.JSON(402, response.Register{
			Basic: response.Basic{
				StatusCode: -1,
				StatusMsg:  "param length err",
			},
			UserId: 1,
			Token:  "zxy",
		})
		return
	}

	if uid, err := service.Register(registerParam.UserName, registerParam.PassWord); err != nil {
		c.JSON(500, response.Register{
			Basic: response.Basic{
				StatusCode: -1,
				StatusMsg:  "register failed",
			},
			UserId: uid,
			Token:  "zxy",
		})
		return
	}

	c.JSON(http.StatusOK, response.Register{
		Basic: response.Basic{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserId: 1,
		Token:  "zxy",
	})
	return
}
