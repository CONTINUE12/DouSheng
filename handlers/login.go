package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	var loginParam request.UserParam

	//绑定参数
	if err := c.ShouldBind(&loginParam); err != nil {
		print("bind param err")
		c.JSON(401, response.Login{
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
	if len(loginParam.UserName) == 0 || len(loginParam.PassWord) == 0 {
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

	var uid int64
	var err error

	if uid, err = service.Login(loginParam.UserName, loginParam.PassWord); err != nil {
		print(err.Error())
		c.JSON(500, response.Login{
			Basic: response.Basic{
				StatusCode: -1,
				StatusMsg:  "login failed" + err.Error(),
			},
			UserId: uid,
			Token:  "zxy",
		})
		return
	}

	token := loginParam.UserName + loginParam.PassWord
	c.JSON(http.StatusOK, response.Login{
		Basic: response.Basic{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserId: uid,
		Token:  token,
	})
}

func DealLogin(c *gin.Context) {

	uidStr := c.Query("user_id")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		print(err.Error())
		return
	}

	if err := service.Verify(uid); err != nil {
		c.JSON(500, response.Login{
			Basic: response.Basic{
				StatusCode: -1,
				StatusMsg:  "verify failed",
			},
		})
		return
	}

	c.JSON(200, response.Login{
		Basic: response.Basic{
			StatusCode: 0,
			StatusMsg:  "verify success",
		},
		UserId: uid,
	})
}
