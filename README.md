# 青训营抖音项目

## 一、小组成员及分工

#### 赵翔宇(组长):负责所有模块代码的编写

#### 王妹湾、杨飞扬、王力、智泽镕:负责下载视频

#### 张佳杰:负责编写汇报ppt

## 二、项目灵感

#### 安卓端的极简版抖音

## 三、项目亮点

#### 1、使用ffmpeg对上传视频进行截图

#### 2、使用雪花算法生成唯一性id，保证id不重复

#### 3、使用ngrok进行内网穿透，让外网能访问到本地资源

## 四、技术说明

#### 本项目使用的技术栈包括：gin(网络框架)、gorm(操作关系型数据库mysql)、ffmpeg(处理上传视频)、ngrok(将内网的8080端口暴露给公网)

## 五、难点突破

#### 1、使用gorm对复杂sql语句(嵌套子查询)进行操作

#### 2、使用了雪花算法生成唯一性id

#### 3、将结构体对象转化为json字符串，交由gin返回

## 六、项目结构

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

## 七、项目成果展示

```shell
https://zxy-dousheng.oss-cn-shenzhen.aliyuncs.com/video/result.mp4
```

