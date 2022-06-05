package repository

import (
	"fmt"
	"time"
)

type Comment struct {
	Id         int64     `gorm:"primary_key"`
	User       User      `gorm:"foreignKey:Id;references:user_id"`
	UserId     int64     `gorm:"column:user_id"`
	VideoId    int64     `gorm:"column:video_id"`
	Content    string    `gorm:"column:content"`
	Token      string    `gorm:"token"`
	CreateDate time.Time `gorm:"create_date"`
}

var ActionComment Comment
var CommentList []Comment

func AddComment(userId, videoId int64, token, commentText string) error {
	comment := &Comment{
		UserId:     userId,
		VideoId:    videoId,
		Content:    commentText,
		Token:      token,
		CreateDate: time.Now(),
	}

	if err := db.Table("comment").Create(comment).Error; err != nil {
		return err
	}

	if err := db.Table("comment").Where("id = ?", comment.Id).Where("user_id IN (?)",
		db.Table("users").Select("id").Where("id = ?", userId)).
		Preload("User").Find(&ActionComment).Error; err != nil {
		return err
	}

	//时间格式要调成MM-dd的格式
	createStr := ActionComment.CreateDate.Format("2006-01-02")
	createDate, _ := time.Parse("2006-01-02", createStr)
	fmt.Printf("%v", createDate)
	return nil
}

func DeleteComment(deleteId int64, token string) error {
	if err := db.Delete(&Comment{}, deleteId).Where("token = ?", token).Error; err != nil {
		return err
	}
	return nil
}

func GetCommentList(videoId int64, token string) error {
	if err := db.Table("comment").Where("token = ?", token).
		Where("video_id = ?", videoId).Preload("User").Find(&CommentList).Error; err != nil {
		return err
	}
	return nil
}
