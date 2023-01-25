package model

import (
	"time"
)
type LikeDB struct {
	UserID 		int64 `gorm:"primaryKey"`
	VideoID 	int64 `gorm:"primaryKey"`
	LikeTime 	time.Time
}

func (*LikeDB) TableName() string {
	return "likes"
}

type LikeDAO struct {}

func NewLikeDAO() LikeDAO {
	return LikeDAO{}
}

func (LikeDAO) CountLikesByUserID(id int64) int64 {
	var count int64
	DB.Model(&LikeDB{}).Where("user_id = ?", id).Count(&count)
	return count
}

func (LikeDAO) CountLikesByVideoID(id int64) int64 {
	var count int64
	DB.Model(&LikeDB{}).Where("video_id = ?", id).Count(&count)
	return count
}

func (LikeDAO) AddNewLike(newLike *LikeDB) {
	DB.Create(newLike)
}

func (LikeDAO) DeleteLike(usr_id int64, video_id int64) {
	DB.Where("user_id = ? AND video_id = ?", usr_id, video_id).Delete(&LikeDB{})
}

