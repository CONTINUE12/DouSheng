package repository

import (
	"sync"
	time2 "time"
)

// Video 需要通过响应返回的Video格式
type Video struct {
	Id            int64      `gorm:"primary_key"`
	Author        User       `gorm:"foreignKey:Id;references:author_id"`
	AuthorId      int64      `sql:"index"`
	PlayUrl       string     `gorm:"column:play_url"`
	CoverUrl      string     `gorm:"column:cover_url"`
	FavoriteCount int64      `gorm:"column:favorite_count"`
	CommentCount  int64      `gorm:"column:comment_count"`
	IsFavorite    bool       `gorm:"column:is_favorite"`
	Title         string     `gorm:"column:title"`
	CreateTime    time2.Time `gorm:"column:create_time"`
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

var LatestList []Video
var PublishList []Video
var NextTime int64

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

// GetLatestVideo 查询最新发布的视频,即发布时间大于给定的最新时间
func (*VideoDao) GetLatestVideo(latestTime int64, token string) error {
	t := time2.Unix(latestTime, 0)
	//从数据库中查找出create_time>latest_time且token为指定的token值的行
	if err := db.Table("videos").
		Where("author_id IN (?)", db.Table("users").
			Select("id").
			Where("token = ?", token)).
		Where("create_time > ?", t).
		Preload("Author").
		Find(&LatestList).Error; err != nil {
		return err
	}

	//将VideoList中最大的create_time作为下次请求的latest_time
	for _, video := range LatestList {
		t := video.CreateTime.Unix()
		if NextTime < t {
			NextTime = t
		}
	}
	return nil
}

// GetUserVideoList 获取用户发布视频的列表
func GetUserVideoList(userId int64, token string) error {
	if err := db.Table("videos").Where("author_id IN (?)",
		db.Table("users").Select("id").Where("id = ?", userId).
			Where("token = ?", token)).Find(&PublishList).Error; err != nil {
		return err
	}
	return nil
}
