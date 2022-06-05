package repository

type Relation struct {
	Id         int64  `gorm:"primary_key"`
	UserId     int64  `gorm:"column:user_id"`
	ToUserId   int64  `gorm:"column:to_user_id"`
	ActionType bool   `gorm:"column:action_type"`
	Token      string `gorm:"column:token"`
}

var FollowUsers []User
var FollowerUsers []User

func Add(userId, toUserId, actionType int64, token string) error {
	var ac bool
	switch actionType {
	case 1:
		ac = true
	case 2:
		ac = false
	}

	relation := &Relation{
		UserId:     userId,
		ToUserId:   toUserId,
		ActionType: ac,
		Token:      token,
	}

	if err := db.Table("relations").Create(relation).Error; err != nil {
		return err
	}
	return nil
}

func FollowList(userId int64, token string) error {
	if err := db.Table("users").Where("token = ?", token).
		Where("id IN (?)",
			db.Table("relations").Select("to_user_id").Where("user_id = ?", userId)).
		Find(&FollowUsers).Error; err != nil {
		return err
	}
	return nil
}

func FollowerList(toUserId int64, token string) error {
	if err := db.Table("users").Where("token = ?", token).
		Where("id IN (?)",
			db.Table("relations").Select("user_id").Where("to_user_id = ?", toUserId)).
		Find(&FollowerUsers).Error; err != nil {
		return err
	}
	return nil
}
