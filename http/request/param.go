package request

//封装发起post请求的参数

// UserParam 登录注册的参数
type UserParam struct {
	UserName string `form:"username" binding:"required"` //此处不能用json,必须要用form,因为post方法的参数是在form里面
	PassWord string `form:"password" binding:"required"`
}

// RelationParam 获取关注列表和粉丝列表的参数
type RelationParam struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

// PublishListParam 获取视频发布列表的参数
type PublishListParam struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}
