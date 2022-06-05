package service

import "douyin/repository"

func AddComment(userId, videoId int64, token, commentText string) error {
	if err := repository.AddComment(userId, videoId, token, commentText); err != nil {
		return err
	}
	return nil
}

func DeleteComment(deleteId int64, token string) error {
	if err := repository.DeleteComment(deleteId, token); err != nil {
		return err
	}
	return nil
}

func GetCommentList(videoId int64, token string) error {
	if err := repository.GetCommentList(videoId, token); err != nil {
		return err
	}
	return nil
}
