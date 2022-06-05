package repository

type Favorite struct {
	Id         int64  `gorm:"primary_key"`
	UserId     int64  `gorm:"column:user_id"`
	VideoId    int64  `gorm:"column:video_id"`
	Token      string `gorm:"column:token"`
	ActionType bool   `gorm:"column:action_type"` //1-点赞,2-取消点赞
}

var FavoriteList []Video

func AddFavorite(userId, videoId, actionType int64, token string) error {
	ac := false
	if actionType == 1 {
		ac = true
	}
	favorite := &Favorite{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: ac,
		Token:      token,
	}

	if err := db.Table("favorite").Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

func GetFavoriteList(userId int64, token string) error {
	if err := db.Table("videos").Where("id IN (?)",
		db.Table("favorite").Select("video_id").
			Where("user_id = ?", userId).
			Where("token = ?", token)).Preload("Author").Find(&FavoriteList).Error; err != nil {
		return err
	}
	return nil
}
