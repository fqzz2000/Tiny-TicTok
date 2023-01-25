package model

import (
	"time"
)
type VideoDB struct {
	VideoID int64 `gorm:"primaryKey"`
	VideoTitle string 
	VideoDesc string
	VideoOwner uint64
	VideoCrtTime time.Time
	VideoFile string
	CoverFile string
}
func (*VideoDB) TableName() string {
	return "videos"
}

type VideoDAO struct {

}
func NewVideoDAO() VideoDAO{
	return VideoDAO{}
}

func (VideoDAO) QueryVideoBeforeTime(t time.Time, limit int) []VideoDB {
	var v []VideoDB;
	DB.Where("video_crt_time > ?", t).Limit(limit).Find(&v);
	return v
}

func (VideoDAO) AddNewVideo(newVideo *VideoDB) {
	DB.Omit("VideoID").Create(newVideo)
}
