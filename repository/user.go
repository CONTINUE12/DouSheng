package repository

import (
	"sync"
)

type User struct {
	Id            int64  `gorm:"primary_key"`
	Name          string `gorm:"column:name"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
	IsFollow      bool   `gorm:"column:is_follow"`
	Token         string `gorm:"column:token"`
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

// CreateUser 创建用户
func (*UserDao) CreateUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		print(err.Error())
		return err
	}
	return nil
}

// QueryUser 查询用户
func (*UserDao) QueryUser(user *User) error {
	if err := db.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// QueryUserId 根据id查询用户
func (*UserDao) QueryUserId(uid int64) error {
	var user User
	if err := db.Model(&User{}).Table("users").Where("id=?", uid).Find(&user).Error; err != nil {
		return err
	}
	return nil
}
