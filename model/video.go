package model

import (
	"time"
)
type VideoDB struct {
	VideoID int64 `gorm:"primaryKey"`
	VideoTitle string 
	VideoDesc string
	VideoOwner int64
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

func (VideoDAO) QueryVideoByID(id int64) VideoDB {
	var v VideoDB
	DB.Where("video_id = ?", id).Find(&v)
	return v
}

func (VideoDAO) QueryVideoBeforeTime(t time.Time, limit int) []VideoDB {
	var v []VideoDB;
	DB.Where("video_crt_time > ?", t).Order("video_crt_time desc").Limit(limit).Find(&v);
	return v
}

func (VideoDAO) QueryVideoByOwnerID(id int64) []VideoDB {
	var v []VideoDB;
	DB.Where("video_owner = ?", id).Order("video_crt_time desc").Find(&v);
	return v
}

func (VideoDAO) CountVideoByOwnerID(id int64) int64 {
	var count int64
	DB.Model(&VideoDB{}).Where("video_owner = ?", id).Count(&count)
	return count
}

func (VideoDAO) AddNewVideo(newVideo *VideoDB) {
	DB.Omit("VideoID").Create(newVideo)
}
