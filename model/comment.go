package model

import (
	"time"
)

type CommentDB struct {
	CommentID 		int64 `gorm:"primaryKey"`
	CommentUserID 	int64 
	CommentVideoID 	int64
	CommentContent 	string
	CommentCrtTime 	time.Time
}

func (*CommentDB) TableName() string {
	return "comments"
}

type CommentDAO struct {}

func NewCommentDAO() CommentDAO {
	return CommentDAO{}
}

func (CommentDAO) CountCommentsByVideoID(id int64) int64 {
	var count int64
	DB.Model(&CommentDB{}).Where("comment_video_id = ?", id).Count(&count)
	return count
}

func (CommentDAO) AddNewComment(newComment *CommentDB) {
	DB.Omit("CommentID").Create(newComment)
}

func (CommentDAO) DeleteComment(id int64) {
	DB.Where("comment_id = ?", id).Delete(&CommentDB{})
}
