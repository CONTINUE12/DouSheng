# 青训营抖音项目

极简版抖音

## 技术栈

#### gin(网络框架)、gorm(操作关系型数据库mysql)、ffmpeg(处理上传视频)、ngrok(将内网的8080端口暴露给公网)

## 项目结构

#### 本项目的树形图如下

```shell
├── README.md       
├── consts                 //保存常用的路径参数
│   └── paths.go
├── go.mod                 //依赖管理
├── handlers               //handler层,接收请求参数和返回响应结果
│   ├── comment_action.go    //添加or删除评论
│   ├── comment_list.go      //获取评论列表
│   ├── favorite_action.go   //点赞or取消点赞
│   ├── favorite_list.go     //获取点赞列表
│   ├── feed.go              //获取视频流信息
│   ├── follow_action.go     //关注or取消关注
│   ├── follow_list.go       //关注列表
│   ├── follower_list.go     //粉丝列表
│   ├── login.go             //登录
│   ├── publish_action.go    //发布视频投稿
│   ├── publish_list.go      //视频投稿列表
│   ├── register.go          //注册
│   └── user_info.go         //登录后获取用户信息
├── http                   //封装http请求和响应的一些参数
│   ├── request 
│   │   └── param.go          //将请求参数封装为结构体
│   └── response
│       └── response.go       //将响应参数封装为结构体
├── main.go                //主启动函数
├── repository             //repository层，面向数据库
│   ├── comment.go             
│   ├── db_init.go
│   ├── favorite.go
│   ├── relation.go
│   ├── user.go
│   └── video.go
├── service                //service层，将handler层获得的数据交给
│   │                        repository层处理
│   ├── comment.go
│   ├── favorite.go
│   ├── relation.go
│   ├── user.go
│   └── video.go
├── utils                  //工具类
│   ├── ffmpeg.go            //ffmpeg，处理上传视频
│   └── snowflake.go         //snowflake，雪花算法，生成唯一性ID
└── youthcamp.sql          //数据库sql文件
```
