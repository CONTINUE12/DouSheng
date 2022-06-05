package service

import "douyin/repository"

func AddFavorite(userId, videoId, actionType int64, token string) error {
	if err := repository.AddFavorite(userId, videoId, actionType, token); err != nil {
		return err
	}
	return nil
}

func GetFavoriteList(userId int64, token string) error {
	if err := repository.GetFavoriteList(userId, token); err != nil {
		return err
	}
	return nil
}
