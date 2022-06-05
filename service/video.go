package service

import "douyin/repository"

// ==============================================获取最新投稿的20个视频===============================================、

func GetLatestVideo(latestTime int64, token string) error {
	if err := repository.NewVideoDaoInstance().GetLatestVideo(latestTime, token); err != nil {
		return err
	}
	return nil
}

// ==============================================获取用户发布的视频列表===========================================

func PublishList(userId int64, token string) error {
	if err := repository.GetUserVideoList(userId, token); err != nil {
		return err
	}
	return nil
}
