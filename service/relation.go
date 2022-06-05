package service

import "douyin/repository"

func Relation(userId, toUserId, actionType int64, token string) error {
	if err := repository.Add(userId, toUserId, actionType, token); err != nil {
		return err
	}
	return nil
}

func FollowList(userId int64, token string) error {
	if err := repository.FollowList(userId, token); err != nil {
		return err
	}
	return nil
}

func FollowerList(toUserId int64, token string) error {
	if err := repository.FollowerList(toUserId, token); err != nil {
		return err
	}
	return nil
}
